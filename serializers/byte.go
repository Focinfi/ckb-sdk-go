package serializers

type Byte byte

func (b Byte) Serialize() []byte {
	return []byte{byte(b)}
}
