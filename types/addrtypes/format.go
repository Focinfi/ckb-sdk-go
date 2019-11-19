package addrtypes

type FormatType = uint8

// Format types
// doc: https://github.com/nervosnetwork/rfcs/blob/master/rfcs/0021-ckb-address-format/0021-ckb-address-format.md#payload-format-types
const (
	FormatTypeShort    FormatType = 1
	FormatTypeFullData FormatType = 2
	FormatTypeFullType FormatType = 4
)

var formatTypeList = []FormatType{FormatTypeShort, FormatTypeFullData, FormatTypeFullType}
var formatTypeFullPayloadList = []FormatType{FormatTypeFullData, FormatTypeFullType}

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
