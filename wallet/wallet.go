package wallet

import (
	"context"
	"fmt"

	"github.com/Focinfi/ckb-sdk-go/address"
	"github.com/Focinfi/ckb-sdk-go/cellcollector"
	"github.com/Focinfi/ckb-sdk-go/key"
	"github.com/Focinfi/ckb-sdk-go/rpc"
	"github.com/Focinfi/ckb-sdk-go/types"
	"github.com/Focinfi/ckb-sdk-go/types/addrtypes"
	"github.com/Focinfi/ckb-sdk-go/types/ckbtypes"
	"github.com/Focinfi/ckb-sdk-go/types/errtypes"
	"github.com/Focinfi/ckb-sdk-go/utils"
)

// Wallet represents a CKB wallet
type Wallet struct {
	Client          *rpc.Client
	Key             *key.Key
	SkipDataAndType bool
	HashType        ckbtypes.HashType
	lockHashHex     *types.HexStr
	lock            *ckbtypes.Script
	sysCells        *utils.SysCells
}

// NewWallet creates and returns a new Wallet, loads the system cells, cache the lock script and compute its hash
func NewWallet(client *rpc.Client, key *key.Key, skipDataAndType bool, hashType ckbtypes.HashType) (*Wallet, error) {
	sysCells, err := utils.LoadSystemCells(*client)
	if err != nil {
		return nil, err
	}
	var codeHash string
	switch hashType {
	case ckbtypes.HashTypeType:
		codeHash = sysCells.Secp256k1TypeHash.Hex()
	case ckbtypes.HashTypeData:
		codeHash = sysCells.Secp256k1CodeHash.Hex()
	default:
		return nil, errtypes.WrapErr(errtypes.WalletErrInvalidHashType, nil)
	}
	lock := ckbtypes.Script{
		Args:     key.Address.KeyHash.Blake160.Hex(),
		CodeHash: codeHash,
		HashType: hashType,
	}
	lockHashHex, err := utils.ScriptHash(lock)
	if err != nil {
		return nil, err
	}

	return &Wallet{
		Client:          client,
		Key:             key,
		SkipDataAndType: skipDataAndType,
		lockHashHex:     lockHashHex,
		sysCells:        sysCells,
		lock:            &lock,
		HashType:        hashType,
	}, nil
}

// NewWalletByPrivKey creates and returns a new Wallet from a private key
func NewWalletByPrivKey(client *rpc.Client, privKey string, skipDataAndType bool, hashType ckbtypes.HashType, mode types.Mode) (*Wallet, error) {
	k, err := key.NewFromPrivKeyHex(privKey, mode)
	if err != nil {
		return nil, err
	}
	return NewWallet(client, k, skipDataAndType, hashType)
}

// Balance gets the balance
func (wallet *Wallet) Balance(ctx context.Context) (uint64, error) {
	collector := cellcollector.NewCellCollector(wallet.Client, wallet.SkipDataAndType)
	_, totalCap, err := collector.GetUnspentCells(ctx, wallet.lockHashHex.Hex(), 0)
	if err != nil {
		return 0, err
	}
	return totalCap, nil
}

// GetUnspentCells get the unspent cells
func (wallet *Wallet) GetUnspentCells(ctx context.Context, needCap uint64) ([]ckbtypes.CellOutputWithOutPoint, error) {
	collector := cellcollector.NewCellCollector(wallet.Client, wallet.SkipDataAndType)
	cells, _, err := collector.GetUnspentCells(ctx, wallet.lockHashHex.Hex(), needCap)
	if err != nil {
		return nil, err
	}
	return cells, nil
}

// GenerateTx generates a transaction ready to send
func (wallet *Wallet) GenerateTx(ctx context.Context, targetAddr string, capacity uint64, data []byte, fee uint64, useDepGroup bool) (*ckbtypes.Transaction, error) {
	config, err := address.Parse(targetAddr, wallet.Key.Address.Mode)
	if err != nil {
		return nil, err
	}
	var codeHash *types.HexStr
	switch config.FormatType {
	case addrtypes.FormatTypeShortLock:
		switch config.CodeHashIndex {
		case addrtypes.CodeHashIndex0:
			codeHash = wallet.sysCells.Secp256k1TypeHash
		case addrtypes.CodeHashIndex1:
			codeHash = wallet.sysCells.MultiSignSecpCellTypeHash
		}
	case addrtypes.FormatTypeCode, addrtypes.FormatTypeData:
		codeHash = config.CodeHash
	}

	dataHex := types.NewHexStr(data)
	output := ckbtypes.Output{
		Capacity: types.HexUint64(capacity).Hex(),
		Lock: ckbtypes.Script{
			Args:     config.Args.Hex(),
			CodeHash: codeHash.Hex(),
			HashType: ckbtypes.HashTypeType,
		},
	}
	outputByteSize, err := output.ByteSize()
	if err != nil {
		return nil, err
	}
	changeOutput := ckbtypes.Output{Lock: *wallet.Lock()}
	minChangeByteSize, err := changeOutput.ByteSize()
	if err != nil {
		return nil, err
	}
	minCap := (outputByteSize + uint64(dataHex.Len())) * types.OneCKBShannon
	minChangeCap := minChangeByteSize * types.OneCKBShannon
	inputs, inputCap, err := wallet.GatherInputs(ctx, capacity, minCap, minChangeCap, fee)
	if err != nil {
		return nil, err
	}
	outputs := []ckbtypes.Output{output}
	outputsData := []string{dataHex.Hex()}
	if changeCap := inputCap - (capacity + fee); changeCap > 0 {
		changeOutput.Capacity = types.HexUint64(changeCap).Hex()
		outputs = append(outputs, changeOutput)
		outputsData = append(outputsData, types.HexStrPrefix)
	}
	witnesses := append([]interface{}{ckbtypes.Witness{}}, utils.EmptyWitnessesByLen(len(inputs)-1)...)
	tx := &ckbtypes.Transaction{
		Version:     types.HexUint64(0).Hex(),
		CellDeps:    []ckbtypes.CellDep{},
		HeaderDeps:  []string{},
		Inputs:      inputs,
		Outputs:     outputs,
		OutputsData: outputsData,
		Witnesses:   witnesses,
	}

	if useDepGroup {
		tx.CellDeps = append(tx.CellDeps, ckbtypes.CellDep{DepType: ckbtypes.DepTypeDepGroup, OutPoint: *wallet.sysCells.Secp256k1GroupOutPoint})
		tx.CellDeps = append(tx.CellDeps, ckbtypes.CellDep{DepType: ckbtypes.DepTypeDepGroup, OutPoint: *wallet.sysCells.MultiSignSecpGroupOutPoint})
	} else {
		tx.CellDeps = append(tx.CellDeps, ckbtypes.CellDep{DepType: ckbtypes.DepTypeCode, OutPoint: *wallet.sysCells.Secp256k1CodeOutPoint})
		tx.CellDeps = append(tx.CellDeps, ckbtypes.CellDep{DepType: ckbtypes.DepTypeCode, OutPoint: *wallet.sysCells.Secp256k1DataOutPoint})
	}

	return utils.SignTransaction(*wallet.Key, *tx)
}

// SendCapacity sends the capacity form the wallet address to the given targetAddr
func (wallet *Wallet) SendCapacity(ctx context.Context, targetAddr string, capacity uint64, data []byte, fee uint64) (string, error) {
	tx, err := wallet.GenerateTx(ctx, targetAddr, capacity, data, fee, true)
	if err != nil {
		return "", err
	}
	return wallet.SendTransaction(ctx, *tx)
}

// DepositToDAO deposit the given capacity into the DAO
func (wallet *Wallet) DepositToDAO(ctx context.Context, capacity, fee uint64) (*ckbtypes.OutPoint, error) {
	sysCells, err := utils.LoadSystemCells(*wallet.Client)
	if err != nil {
		return nil, err
	}
	output := ckbtypes.Output{
		Capacity: types.HexUint64(capacity).Hex(),
		Lock:     *wallet.Lock(),
		Type: &ckbtypes.Script{
			Args:     types.HexStrPrefix,
			CodeHash: sysCells.DaoTypeHash.Hex(),
			HashType: ckbtypes.HashTypeType,
		},
	}
	outputByteSize, err := output.ByteSize()
	if err != nil {
		return nil, err
	}
	changeOutput := ckbtypes.Output{
		Lock: *wallet.Lock(),
	}
	changeOutputByteSize, err := changeOutput.ByteSize()
	if err != nil {
		return nil, err
	}
	minCap := (outputByteSize + uint64(daoDepositOutputDataHex.Len())) * types.OneCKBShannon
	minChangeCap := changeOutputByteSize * types.OneCKBShannon
	inputs, inputsCap, err := wallet.GatherInputs(ctx, capacity, minCap, minChangeCap, fee)
	if err != nil {
		return nil, err
	}
	outputs := []ckbtypes.Output{output}
	outputsData := []string{daoDepositOutputDataHex.Hex()}
	changeOutputCap := inputsCap - (capacity + fee)
	if changeOutputCap > 0 {
		changeOutput.Capacity = types.HexUint64(changeOutputCap).Hex()
		outputs = append(outputs, changeOutput)
		outputsData = append(outputsData, types.HexStrPrefix)
	}
	witnesses := append([]interface{}{ckbtypes.Witness{}}, utils.EmptyWitnessesByLen(len(inputs)-1)...)
	tx := ckbtypes.Transaction{
		Version:    types.HexUint64(0).Hex(),
		HeaderDeps: []string{},
		CellDeps: []ckbtypes.CellDep{
			{OutPoint: *wallet.sysCells.Secp256k1GroupOutPoint, DepType: ckbtypes.DepTypeDepGroup},
			{OutPoint: *wallet.sysCells.DaoOutPoint, DepType: ckbtypes.DepTypeCode},
		},
		Inputs:      inputs,
		Outputs:     outputs,
		OutputsData: outputsData,
		Witnesses:   witnesses,
	}
	signedTx, err := utils.SignTransaction(*wallet.Key, tx)
	if err != nil {
		return nil, err
	}
	txHash, err := wallet.SendTransaction(ctx, *signedTx)
	if err != nil {
		return nil, err
	}
	return &ckbtypes.OutPoint{
		TxHash: txHash,
		Index:  types.HexUint64(0).Hex(),
	}, nil
}

// StartWithdrawingFromDAO sends the transaction to create the withdraw cell form the deposit cell
func (wallet *Wallet) StartWithdrawingFromDAO(ctx context.Context, depositOutPoint ckbtypes.OutPoint, fee uint64) (*ckbtypes.OutPoint, error) {
	cellInfo, err := wallet.Client.GetLiveCell(ctx, depositOutPoint, false)
	if err != nil {
		return nil, err
	}
	if cellInfo.Status != ckbtypes.CellStatusLive {
		return nil, errtypes.WrapErr(errtypes.DAOWithdrawErrDepositCellNotLive, nil)
	}
	txInfo, err := wallet.Client.GetTransaction(ctx, depositOutPoint.TxHash)
	if err != nil {
		return nil, err
	}
	if txInfo.Status.Status != ckbtypes.TransactionStatusCommitted {
		return nil, errtypes.WrapErr(errtypes.DAOWithdrawErrDepositTxNotCommitted, nil)
	}
	depositBlock, err := wallet.Client.GetBlock(ctx, txInfo.Status.BlockHash)
	if err != nil {
		return nil, err
	}
	depositBlockNum := depositBlock.Header.Number
	depositBlockNumHex, err := types.ParseHexUint64(depositBlockNum)
	if err != nil {
		return nil, err
	}
	depositBlockNumHexStr := types.NewHexStr(depositBlockNumHex.LittleEndianBytes(8))

	output := cellInfo.Cell.Output.Clone()
	outputData := depositBlockNumHexStr.Hex()

	changeOutput := ckbtypes.Output{Lock: *wallet.Lock()}
	minChangeCap, err := changeOutput.ByteSize()
	if err != nil {
		return nil, err
	}

	inputs, inputsCap, err := wallet.GatherInputs(ctx, 0, 0, minChangeCap, fee)
	if err != nil {
		return nil, err
	}
	outputs := []ckbtypes.Output{*output}
	outputsData := []string{outputData}
	changeCap := inputsCap - fee
	if changeCap > 0 {
		changeOutput.Capacity = types.HexUint64(changeCap).Hex()
		outputs = append(outputs, changeOutput)
		outputsData = append(outputsData, types.HexStrPrefix)
	}

	firstInput := ckbtypes.Input{
		PreviousOutput: *depositOutPoint.Clone(),
		Since:          types.HexUint64(0).Hex(),
	}
	inputs = append([]ckbtypes.Input{firstInput}, inputs...)
	witness := append([]interface{}{ckbtypes.Witness{}}, utils.EmptyWitnessesByLen(len(inputs)-1)...)
	tx := ckbtypes.Transaction{
		Version: types.HexUint64(0).Hex(),
		CellDeps: []ckbtypes.CellDep{
			{OutPoint: *wallet.sysCells.Secp256k1GroupOutPoint, DepType: ckbtypes.DepTypeDepGroup},
			{OutPoint: *wallet.sysCells.DaoOutPoint, DepType: ckbtypes.DepTypeCode},
		},
		HeaderDeps:  []string{depositBlock.Header.Hash},
		Inputs:      inputs,
		Outputs:     outputs,
		OutputsData: outputsData,
		Witnesses:   witness,
	}
	signedTx, err := utils.SignTransaction(*wallet.Key, tx)
	if err != nil {
		return nil, err
	}
	txHash, err := wallet.SendTransaction(ctx, *signedTx)
	if err != nil {
		return nil, err
	}
	return &ckbtypes.OutPoint{
		TxHash: txHash,
		Index:  types.HexUint64(0).Hex(),
	}, nil
}

// GenWithdrawFromDAOTx generates the transaction to withdraw from the DAO
func (wallet *Wallet) GenWithdrawFromDAOTx(ctx context.Context, depositOutpoint, withdrawingOutpoint ckbtypes.OutPoint, fee uint64) (*ckbtypes.Transaction, error) {
	cellInfo, err := wallet.Client.GetLiveCell(ctx, withdrawingOutpoint, true)
	if err != nil {
		return nil, err
	}
	if cellInfo.Status != ckbtypes.CellStatusLive {
		return nil, errtypes.WrapErr(errtypes.DAOWithdrawErrDepositCellNotLive, nil)
	}
	txInfo, err := wallet.Client.GetTransaction(ctx, withdrawingOutpoint.TxHash)
	if err != nil {
		return nil, err
	}
	if txInfo.Status.Status != ckbtypes.TransactionStatusCommitted {
		return nil, errtypes.WrapErr(errtypes.DAOWithdrawErrDepositTxNotCommitted, nil)
	}

	depositBlockNumberHex, err := types.ParseHexUint64(cellInfo.Cell.Data.Content)
	if err != nil {
		return nil, err
	}
	reversedDepositNumberHex := types.NewHexStr(depositBlockNumberHex.LittleEndianBytes(8)).Hex()
	depositBlockNumberHex, err = types.ParseHexUint64(reversedDepositNumberHex)
	if err != nil {
		return nil, err
	}
	depositBlock, err := wallet.Client.GetBlockByNumber(ctx, *depositBlockNumberHex)
	if err != nil {
		return nil, err
	}
	depositEpoch, err := utils.ParseEpochByHexStr(depositBlock.Header.Epoch)
	if err != nil {
		return nil, err
	}

	withdrawBlock, err := wallet.Client.GetBlock(ctx, txInfo.Status.BlockHash)
	if err != nil {
		return nil, err
	}
	withdrawEpoch, err := utils.ParseEpochByHexStr(withdrawBlock.Header.Epoch)
	if err != nil {
		return nil, err
	}

	withdrawFraction := withdrawEpoch.Index * depositEpoch.Length
	depositFraction := depositEpoch.Index * withdrawEpoch.Length
	depositEpochs := withdrawEpoch.Number - depositEpoch.Number
	if withdrawFraction > depositFraction {
		depositEpochs += 1
	}
	lockEpochs := (depositEpochs + (DAOLockPeriodEpochs - 1)) / DAOLockPeriodEpochs * DAOLockPeriodEpochs
	minimalEpoch := utils.Epoch{
		Number: depositEpoch.Number + lockEpochs,
		Index:  depositEpoch.Index,
		Length: depositEpoch.Length,
	}
	minimalSince := minimalEpoch.Since()
	outputCapacity, err := wallet.Client.CalculateDAOMaximumWithdraw(ctx, *depositOutpoint.Clone(), withdrawBlock.Header.Hash)
	outputs := []ckbtypes.Output{
		{
			Capacity: types.HexUint64(outputCapacity - fee).Hex(),
			Lock:     *wallet.Lock(),
		},
	}
	outputsData := []string{types.HexStrPrefix}
	tx := ckbtypes.Transaction{
		Version: types.HexUint64(0).Hex(),
		CellDeps: []ckbtypes.CellDep{
			{
				OutPoint: *wallet.sysCells.DaoOutPoint,
				DepType:  ckbtypes.DepTypeCode,
			},
			{
				OutPoint: *wallet.sysCells.Secp256k1GroupOutPoint,
				DepType:  ckbtypes.DepTypeDepGroup,
			},
		},
		HeaderDeps: []string{
			depositBlock.Header.Hash,
			withdrawBlock.Header.Hash,
		},
		Inputs: []ckbtypes.Input{
			{
				PreviousOutput: *withdrawingOutpoint.Clone(),
				Since:          types.HexUint64(minimalSince).Hex(),
			},
		},
		Outputs:     outputs,
		OutputsData: outputsData,
		Witnesses: []interface{}{
			ckbtypes.Witness{InputType: "0x0000000000000000"},
		},
	}
	return utils.SignTransaction(*wallet.Key, tx)
}

// GetTransaction gets the transaction
func (wallet *Wallet) GetTransaction(ctx context.Context, hash string) (*ckbtypes.TransactionInfo, error) {
	return wallet.Client.GetTransaction(ctx, hash)
}

// BlockAssemblerConfig gets the block assembler config used in ckb.toml
func (wallet *Wallet) BlockAssemblerConfig() string {
	return fmt.Sprintf(
		`[block_assembler]
	code_hash = %s
	args = %s`,
		wallet.Lock().CodeHash,
		wallet.Lock().Args)
}

// SendTransaction sends the transaction
func (wallet *Wallet) SendTransaction(ctx context.Context, transaction ckbtypes.Transaction) (string, error) {
	return wallet.Client.SendTransaction(ctx, transaction.ToRaw())
}

// GatherInputs gathers inputs
func (wallet *Wallet) GatherInputs(ctx context.Context, capacity, minCap, minChangeCap, fee uint64) ([]ckbtypes.Input, uint64, error) {
	collector := cellcollector.NewCellCollector(wallet.Client, wallet.SkipDataAndType)
	return collector.GatherInputs(ctx, []string{wallet.lockHashHex.Hex()}, capacity, minCap, minChangeCap, fee)
}

// Lock gets the lock of this wallet
func (wallet *Wallet) Lock() *ckbtypes.Script {
	return wallet.lock
}

// Lock gets the lock hash
func (wallet *Wallet) CodeHash(ctx context.Context) string {
	return wallet.Lock().CodeHash
}
