package serializers

type Witness = ByteFixVec

type Witnesses = ByteDynVec

func NewWitnessesByHexes(hexes []string) (*ByteDynVec, error) {
	return NewByteDynVecByHexes(hexes)
}
