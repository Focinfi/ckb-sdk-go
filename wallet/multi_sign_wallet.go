package wallet

import (
	"context"
	"errors"

	"github.com/Focinfi/ckb-sdk-go/address"
	"github.com/Focinfi/ckb-sdk-go/cellcollector"
	"github.com/Focinfi/ckb-sdk-go/crypto/blake160"
	"github.com/Focinfi/ckb-sdk-go/rpc"
	"github.com/Focinfi/ckb-sdk-go/types"
	"github.com/Focinfi/ckb-sdk-go/types/addrtypes"
	"github.com/Focinfi/ckb-sdk-go/types/ckbtypes"
	"github.com/Focinfi/ckb-sdk-go/types/errtypes"
	"github.com/Focinfi/ckb-sdk-go/utils"
)

type MultiSignWalletConfig struct {
	RequireN  uint8
	Threshold uint8
	PubKeys   []string
	Since     uint64
}

func NewMultiSignWalletConfig(requireN uint8, threshold uint8, pubKeys []string, since uint64) (*MultiSignWalletConfig, error) {
	if len(pubKeys) > 255 {
		return nil, errtypes.WrapErr(errtypes.MultiSignWalletConfigErrPubKeysNumberTooBig, errors.New("must less than 256"))
	}
	return &MultiSignWalletConfig{
		RequireN:  requireN,
		Threshold: threshold,
		PubKeys:   pubKeys,
		Since:     since,
	}, nil
}

func (config MultiSignWalletConfig) Serialize() ([]byte, error) {
	meta := []byte{
		config.RequireN,
		config.Threshold,
		uint8(len(config.PubKeys)),
	}
	body := make([]byte, 0, len(config.PubKeys)*20)
	for _, pubKey := range config.PubKeys {
		h, err := blake160.Blake160(pubKey)
		if err != nil {
			return nil, err
		}
		body = append(body, h...)
	}
	return append(meta, body...), nil
}

func (config MultiSignWalletConfig) Blake160() ([]byte, error) {
	bin, err := config.Serialize()
	if err != nil {
		return nil, err
	}
	return blake160.Blake160Binary(bin)
}

func (config MultiSignWalletConfig) LockArgs() (*types.HexStr, error) {
	blake160Bytes, err := config.Blake160()
	if err != nil {
		return nil, err
	}
	sinceBytes := types.HexUint64(config.Since).LittleEndianBytes(8)
	return types.NewHexStr(blake160Bytes).AppendBytes(sinceBytes), nil
}

type MultiSignWallet struct {
	Client          *rpc.Client
	Config          *MultiSignWalletConfig
	SkipDataAndType bool
	Mode            types.Mode
	sysCells        *utils.SysCells
	lock            *ckbtypes.Script
	lockHash        *types.HexStr
}

func NewMultiSignWallet(client rpc.Client, config MultiSignWalletConfig, skipDataAndType bool, mode types.Mode) (*MultiSignWallet, error) {
	sysCells, err := utils.LoadSystemCells(client)
	if err != nil {
		return nil, err
	}
	lockArgs, err := config.LockArgs()
	if err != nil {
		return nil, err
	}
	lock := ckbtypes.Script{
		Args:     lockArgs.Hex(),
		CodeHash: sysCells.MultiSignSecpCellTypeHash.Hex(),
		HashType: ckbtypes.HashTypeType,
	}
	lockHash, err := utils.ScriptHash(lock)
	if err != nil {
		return nil, err
	}

	return &MultiSignWallet{
		Client:          &client,
		Config:          &config,
		SkipDataAndType: skipDataAndType,
		Mode:            mode,
		sysCells:        sysCells,
		lock:            &lock,
		lockHash:        lockHash,
	}, nil
}

func (wallet MultiSignWallet) Address() (string, error) {
	lockArgs, err := wallet.Config.LockArgs()
	if err != nil {
		return "", err
	}
	payload := append([]byte{addrtypes.FormatTypeFullType})
	payload = append(payload, wallet.sysCells.MultiSignSecpCellTypeHash.Bytes()...)
	payload = append(payload, lockArgs.Bytes()...)
	return address.EncodeAddress(addrtypes.PrefixForMode(wallet.Mode), payload)
}

func (wallet MultiSignWallet) GetBalance(ctx context.Context) (uint64, error) {
	collector := cellcollector.NewCellCollector(wallet.Client, wallet.SkipDataAndType)
	_, capacity, err := collector.GetUnspentCells(ctx, wallet.lockHash.Hex(), 0)
	return capacity, err
}

func (wallet MultiSignWallet) Lock() ckbtypes.Script {
	return *wallet.lock
}

func (wallet MultiSignWallet) GenerateTx(ctx context.Context, targetAddr string, capacity uint64, privKeys []string, data []byte, fee uint64) (*ckbtypes.Transaction, error) {
	if len(privKeys) != int(wallet.Config.Threshold) {
		return nil, errtypes.WrapErr(errtypes.MultiSignWalletPrivKeysNumberNotMatchThreshold, nil)
	}
	targetAddrLock, err := utils.LockScriptFormAddress(targetAddr, wallet.Mode, *wallet.sysCells)
	if err != nil {
		return nil, err
	}

	output := ckbtypes.Output{
		Capacity: types.HexUint64(capacity).Hex(),
		Lock:     *targetAddrLock,
	}

	changeOutput := ckbtypes.Output{Lock: *wallet.lock}
	outputCap, err := output.ByteSize()
	if err != nil {
		return nil, err
	}
	changeOutputCap, err := changeOutput.ByteSize()
	if err != nil {
		return nil, err
	}
	minCap := (outputCap + uint64(len(data))) * types.OneCKBShannon
	minChangeCap := changeOutputCap * types.OneCKBShannon
	collector := cellcollector.NewCellCollector(wallet.Client, wallet.SkipDataAndType)
	inputs, inputsCap, err := collector.GatherInputs(ctx, []string{wallet.lockHash.Hex()}, capacity, minCap, minChangeCap, fee)
	if err != nil {
		return nil, err
	}

	outputs := []ckbtypes.Output{output}
	outputsData := []string{types.NewHexStr(data).Hex()}

	changeCap := inputsCap - (capacity + fee)
	if changeCap > 0 {
		changeOutput.Capacity = types.HexUint64(changeCap).Hex()
		outputs = append(outputs, changeOutput)
		outputsData = append(outputsData, types.HexStrPrefix)
	}

	for _, input := range inputs {
		input.Since = types.HexUint64(wallet.Config.Since).Hex()
	}
	witnesses := append([]interface{}{ckbtypes.Witness{}}, utils.EmptyWitnessesByLen(len(inputs)-1)...)
	tx := ckbtypes.Transaction{
		Version: types.Hex0.Hex(),
		CellDeps: []ckbtypes.CellDep{
			{OutPoint: *wallet.sysCells.MultiSignSecpGroupOutPoint, DepType: ckbtypes.DepTypeDepGroup},
		},
		HeaderDeps:  []string{},
		Inputs:      inputs,
		Outputs:     outputs,
		OutputsData: outputsData,
		Witnesses:   witnesses,
	}

	configSerialization, err := wallet.Config.Serialize()
	if err != nil {
		return nil, err
	}
	return utils.MultiSignTransaction(privKeys, tx, types.NewHexStr(configSerialization), wallet.Mode)
}

func (wallet MultiSignWallet) SendCapacity(ctx context.Context, targetAddr string, capacity uint64, privKeys []string, data []byte, fee uint64) (string, error) {
	transaction, err := wallet.GenerateTx(ctx, targetAddr, capacity, privKeys, data, fee)
	if err != nil {
		return "", err
	}
	return wallet.Client.SendTransaction(ctx, transaction.ToRaw())
}
