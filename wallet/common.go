package wallet

import "github.com/Focinfi/ckb-sdk-go/types"

// daoDepositOutputDataHex 0x0000000000000000
var daoDepositOutputDataHex = types.NewHexStr(make([]byte, 8))

const (
	DAOLockPeriodEpochs = 180
	DAOMaturityBlocks   = 5
)
