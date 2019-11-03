package serializers

type Serializer interface {
	Serialize() []byte
}

const (
	Uint32Capacity         = 4
	Uint64Capacity         = 8
	Byte32Capacity         = 32
	ScriptHashTypeCapacity = 1
	OutPointCapacity       = Uint64Capacity + Byte32Capacity
	InputCapacity          = OutPointCapacity + Uint64Capacity
)

type (
	Capacity       = Uint64
	Since          = Uint64
	CodeHash       = Byte32
	DepType        = Byte
	HashType       = Byte
	HeaderDep      = Byte32
	OutPointTxHash = Byte32
	OutPointIndex  = Uint32
	Arg            = ByteFixVec
	OutputData     = ByteFixVec
	Witness        = ByteFixVec
	Version        = Uint32
)

var (
	DepTypeCode     = DepType(0)
	DepTypeDepGroup = DepType(1)
	HashTypeType    = HashType(1)
	HashTypeData    = HashType(0)
)
