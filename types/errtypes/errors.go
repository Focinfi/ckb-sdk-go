package errtypes

import "fmt"

type ErrType string

// rpc error
const (
	RPCErrMarshalRequestBodyFail    ErrType = "rpc-marshal-request-body-error"
	RPCErrNewRequestFail            ErrType = "rpc-new-request-error"
	RPCErrHTTPRequestFail           ErrType = "rpc-http-request-error"
	RPCErrReadResponseBodyFail      ErrType = "rpc-read-response-body-error"
	RPCErrUnmarshalResponseBodyFail ErrType = "rpc-unmarshal-response-body-error"
	RPCErrResultErrorFail           ErrType = "rpc-result-fail-error"
	RPCErrResponseDataBroken        ErrType = "rpc-response-data-broken-error"
	RPCErrGetInputCapacityTooSmall  ErrType = "rpc-get-input-capacity-too-small-error"
	RPCErrGetInputCapacityNotEnough ErrType = "rpc-input-capacity-not-enough-error"
	RPCErrGetGenesisBlockBroken     ErrType = "rpc-get-genesis-block-broken-error"
)

// address error
const (
	AddressErrEmptyPubKey          ErrType = "address-empty-public-key-error"
	AddressErrConvertBitFail       ErrType = "address-convert-bit-error"
	AddressErrInvalidPrefix        ErrType = "address-invalid-prefix-error"
	AddressErrTooShort             ErrType = "address-too-short-error"
	AddressErrInvalidHashTypeIndex ErrType = "address-invalid-hash-type-index-error"
	AddressErrFormatTypeWrong      ErrType = "address-format-type-wrong-error"
)

// wallet
const (
	WalletErrInvalidHashType                       ErrType = "wallet-invalid-hash-type-err"
	MultiSignWalletConfigErrPubKeysNumberTooBig    ErrType = "multi-sign-wallet-config-public-keys-number-too-big-error"
	MultiSignWalletPrivKeysNumberNotMatchThreshold ErrType = "multi-sign-address-private-keys-number-match-threshold-error"
)

// key error
const (
	KeyErrPrivateKeySizeWrong ErrType = "key-private-key-wrong-size-error"
)

// hex error
const (
	HexErrNeed0xPrefix   ErrType = "hex-need-0x-prefix-error"
	HexErrStrFormatWrong ErrType = "hex-wrong-format-error"
)

// crypto error
const (
	CryptoErrBlake160Fail        ErrType = "crypto-blake160-error"
	CryptoErrBech32EncodeFail    ErrType = "crypto-bech32-encode-error"
	CryptoErrBech32DecodeFail    ErrType = "crypto-bech32-decode-error"
	CryptoErrSignDataFail        ErrType = "crypto-sign-data-error"
	CryptoErrDataByteCountNot32  ErrType = "crypto-data-byte-count-not-32-error"
	CryptoErrSignRecoverableFail ErrType = "crypto-sign-recoverable-error"
)

// serialization error
const (
	SerializationErrItemSizeNotFixed              ErrType = "serialization-item-size-not-fixed-error"
	SerializationErrInvalidVersion                ErrType = "serialization-invalid-version-error"
	SerializationErrUnknownDepType                ErrType = "serialization-unknown-dep-type-error"
	SerializationErrInvalidCellDep                ErrType = "serialization-invalid-cell-dep-error"
	SerializationErrInvalidOutPointIndex          ErrType = "serialization-invalid-out-point-index-error"
	SerializationErrInvalidOutPointHash           ErrType = "serialization-invalid-out-point-hash-error"
	SerializationErrByte32WrongLen                ErrType = "serialization-byte32-wrong-len-error"
	SerializationErrInvalidHeaderDep              ErrType = "serialization-invalid-header-dep-error"
	SerializationErrInvalidInput                  ErrType = "serialization-invalid-input-error"
	SerializationErrInvalidSince                  ErrType = "serialization-invalid-input-since-error"
	SerializationErrOutputCapacity                ErrType = "serialization-invalid-output-capacity-error"
	SerializationErrInvalidCodeHash               ErrType = "serialization-invalid-code-hash-error"
	SerializationErrInvalidHashType               ErrType = "serialization-invalid-hash-type-error"
	SerializationErrInvalidArgs                   ErrType = "serialization-invalid-args-error"
	SerializationErrInvalidLockScript             ErrType = "serialization-invalid-lock-script-error"
	SerializationErrInvalidTypeScript             ErrType = "serialization-invalid-lock-type-error"
	SerializationErrInvalidOutput                 ErrType = "serialization-invalid-output-error"
	SerializationErrInvalidOutputData             ErrType = "serialization-invalid-output-data-error"
	SerializationErrInvalidWitness                ErrType = "serialization-invalid-witness-error"
	SerializationErrInvalidWitnessesForInputLock  ErrType = "serialization-invalid-witness-for-input-lock-error"
	SerializationErrInvalidWitnessesForInputType  ErrType = "serialization-invalid-witness-for-input-type-error"
	SerializationErrInvalidWitnessesForOutputType ErrType = "serialization-invalid-witness-for-output-type-error"
)

// generate transaction error
const (
	GenTransErrGetOutputLockByteSizeFail ErrType = "gen-trans-get-output-lock-byte-size-error"
	GenTransErrGetOutputTypeByteSizeFail ErrType = "gen-trans-get-output-type-byte-size-error"
	GenTransErrWitnessNotEnough          ErrType = "gen-trans-witness-not-enough-error"
	GenTransErrFirstWitnessTypeWrong     ErrType = "gen-trans-first-witness-format-wrong-error"
	GenTransErrHexWitnessTypeWrong       ErrType = "gen-trans-hex-witness-type-wrong-error"
	GenTransErrSignFail                  ErrType = "gen-trans-sign-fail-error"
)

// dao error
const (
	DAOWithdrawErrDepositCellNotLive    ErrType = "dao-withdraw-deposit-cell-not-live-error"
	DAOWithdrawErrDepositTxNotCommitted ErrType = "dao-withdraw-deposit-cell-tx-not-committed-error"
)

type Error struct {
	ErrType ErrType
	Message string
}

func WrapErr(errType ErrType, e error) Error {
	message := ""
	if e != nil {
		message = e.Error()
	}
	return Error{
		ErrType: errType,
		Message: message,
	}
}

func (err Error) IsA(errType ErrType) bool {
	return err.ErrType == errType
}

func (err Error) Error() string {
	return fmt.Sprintf("[%s] %s", err.ErrType, err.Message)
}
