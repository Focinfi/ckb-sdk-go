package wallet

import (
	"context"

	"github.com/Focinfi/ckb-sdk-go/utils"

	"github.com/Focinfi/ckb-sdk-go/cellcollector"

	"github.com/Focinfi/ckb-sdk-go/key"
	"github.com/Focinfi/ckb-sdk-go/rpc"
	"github.com/Focinfi/ckb-sdk-go/types"
	"github.com/Focinfi/ckb-sdk-go/types/ckbtypes"
)

type Wallet struct {
	Client          *rpc.Client
	Key             *key.Key
	SkipDataAndType bool
	lockHashHex     *types.HexStr
	lock            *ckbtypes.Script
}

func NewWallet(client *rpc.Client, key *key.Key, skipDataAndType bool) (*Wallet, error) {
	lockHashHex, err := utils.LockScriptHash(key.Address.PubKey.PubKey)
	if err != nil {
		return nil, err
	}

	return &Wallet{
		Client:          client,
		Key:             key,
		SkipDataAndType: skipDataAndType,
		lockHashHex:     lockHashHex,
	}, nil
}

func NewWalletByPrivKey(client *rpc.Client, privKey string, skipDataAndType bool, mode types.Mode) (*Wallet, error) {
	key, err := key.NewFromPrivKeyHex(privKey, mode)
	if err != nil {
		return nil, err
	}
	return NewWallet(client, key, skipDataAndType)
}

func (wallet *Wallet) Balance(ctx context.Context) (uint64, error) {
	collector := cellcollector.NewCellCollector(wallet.Client, wallet.SkipDataAndType)
	_, totalCap, err := collector.GetUnspentCells(ctx, wallet.lockHashHex.Hex(), 0)
	if err != nil {
		return 0, err
	}
	return totalCap, nil
}

func (wallet *Wallet) GetUnspentCells(ctx context.Context, needCap uint64) ([]ckbtypes.Cell, error) {
	collector := cellcollector.NewCellCollector(wallet.Client, wallet.SkipDataAndType)
	cells, _, err := collector.GetUnspentCells(ctx, wallet.lockHashHex.Hex(), needCap)
	if err != nil {
		return nil, err
	}
	return cells, nil
}

func (wallet *Wallet) GenerateTx(ctx context.Context, targetAddr string, capacity uint64, data []byte, fee uint64, useDepGroup bool) (*ckbtypes.Transaction, error) {
	return nil, nil
}

func (wallet *Wallet) SendCapacity(ctx context.Context, targetAddr string, capacity uint64, data []byte, fee uint64) {
}

func (wallet *Wallet) DepositToDAO(ctx context.Context, capacity, fee uint64) (*ckbtypes.OutPoint, error) {
	return nil, nil
}

func (wallet *Wallet) GenerateWithdrawFromDAOTransaction(ctx context.Context, point ckbtypes.OutPoint, fee uint64) (*ckbtypes.Transaction, error) {
	return nil, nil
}

func (wallet *Wallet) GetTransaction(ctx context.Context, hash string) (*ckbtypes.Transaction, error) {
	return nil, nil
}

func (wallet *Wallet) BlockAssemblerConfig() string {
	//	return fmt.Sprintf(
	//		`[block_assembler]
	//code_hash = %s
	//args = %s`)
	return ""
}

func (wallet *Wallet) SendTransaction(ctx context.Context, transaction ckbtypes.Transaction) (string, error) {
	return "", nil
}

func (wallet *Wallet) GatherInputs(ctx context.Context, capacity, minCap, minChangeCap, fee uint64) ([]ckbtypes.Input, error) {
	return nil, nil
}

func (wallet *Wallet) Lock() ckbtypes.Script {
	if wallet.lock == nil {
		// TODO: init lock
	}
	return *wallet.lock
}

func (wallet *Wallet) CodeHash(ctx context.Context) (string, error) {
	return "", nil
}
