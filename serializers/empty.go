package serializers

type Empty struct{}

func (Empty) Serialize() []byte {
	return []byte{}
}
