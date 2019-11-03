package serializers

type Array []Serializer

func (arr Array) Serialize() []byte {
	bs := make([]byte, 0)
	for _, slr := range arr {
		bs = append(bs, slr.Serialize()...)
	}
	return bs
}
