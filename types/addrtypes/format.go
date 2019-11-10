package addrtypes

type FormatType = uint8

// Format types
// doc: https://github.com/nervosnetwork/rfcs/blob/master/rfcs/0021-ckb-address-format/0021-ckb-address-format.md#payload-format-types
const (
	FormatTypeShortLock FormatType = 1
	FormatTypeData      FormatType = 2
	FormatTypeCode      FormatType = 4
)

var formatTypeList = []FormatType{FormatTypeShortLock, FormatTypeData, FormatTypeCode}
var formatTypeFullPayloadList = []FormatType{FormatTypeShortLock, FormatTypeData, FormatTypeCode}

func IsAllowdedFormatType(formatType FormatType) bool {
	for _, ft := range formatTypeList {
		if formatType == ft {
			return true
		}
	}
	return false
}

func IsFullPayloadFormatType(formatType FormatType) bool {
	for _, ft := range formatTypeFullPayloadList {
		if formatType == ft {
			return true
		}
	}
	return false
}
