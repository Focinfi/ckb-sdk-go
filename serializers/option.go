package serializers

type ByteFixVecOption struct {
	serializer Serializer
}

func (option ByteFixVecOption) Serialize() []byte {
	if option.serializer != nil {
		return option.serializer.Serialize()
	}
	return nil
}

func NewByteFixVecOption(hexStr string) (vec *ByteFixVecOption, err error) {
	var serializer Serializer
	if hexStr == "" {
		serializer = Empty{}
	} else {
		serializer, err = NewByteFixVecByHex(hexStr)
		if err != nil {
			return nil, err
		}
	}
	return &ByteFixVecOption{
		serializer: serializer,
	}, nil
}
