package serializers

type DynVec struct {
	FullLen Uint32
	Offsets []Uint32
	data    []byte
}

func (dv *DynVec) Serialize() []byte {
	return append(dv.Header(), dv.data...)
}

func (dv *DynVec) Header() []byte {
	bs := make([]byte, 0, 4*(1+len(dv.Offsets)))
	bs = append(bs, dv.FullLen.Serialize()...)
	for _, offset := range dv.Offsets {
		bs = append(bs, offset.Serialize()...)
	}
	return bs
}

func NewDynVec(fields []Serializer) *DynVec {
	data := make([]byte, 0)
	capacities := make([]int, 0, len(fields))
	for _, field := range fields {
		bs := field.Serialize()
		data = append(data, bs...)
		capacities = append(capacities, len(bs))
	}

	offsets := make([]Uint32, 0, len(fields))
	lastOffset := Uint32(Uint32Capacity * (1 + len(fields)))
	if len(fields) > 0 {
		for i := 0; i < len(fields); i++ {
			offsets = append(offsets, lastOffset)
			lastOffset += Uint32(capacities[i])
		}
	}

	fullBytesSize := 4 + len(offsets)*4 + len(data)
	table := &DynVec{
		FullLen: Uint32(fullBytesSize),
		Offsets: offsets,
		data:    data,
	}
	return table
}

// ByteDynVec dyn serializers [fix serializers]
type ByteDynVec struct {
	serializer *DynVec
}

func (vec ByteDynVec) Serialize() []byte {
	return vec.serializer.Serialize()
}

func NewByteDynVecByHexes(hexes []string) (*ByteDynVec, error) {
	byteFixVecs := make([]Serializer, 0, len(hexes))
	for _, hex := range hexes {
		fixVec, err := NewByteFixVecByHex(hex)
		if err != nil {
			return nil, err
		}
		byteFixVecs = append(byteFixVecs, fixVec)
	}
	return &ByteDynVec{
		serializer: NewDynVec(byteFixVecs),
	}, nil
}
