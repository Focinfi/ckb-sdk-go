package serializers

type Table struct {
	serializer *DynVec
}

func (dv Table) Serialize() []byte {
	return dv.serializer.Serialize()
}

func NewTable(fields []Serializer) *Table {
	return &Table{
		serializer: NewDynVec(fields),
	}
}
