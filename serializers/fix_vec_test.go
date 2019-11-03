package serializers

import "testing"

func TestNewByteFixVecByHex(t *testing.T) {
	vec := NewByteFixVec([]byte("abcd0x1234"))
	t.Logf("%x", vec.Serialize())
}
