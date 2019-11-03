package serializers

type Struct []Serializer

func (arr Struct) Serialize() []byte {
	bs := make([]byte, 0)
	for _, slr := range arr {
		bs = append(bs, slr.Serialize()...)
	}
	return bs
}
