package types

import "testing"

func TestHexNumber(t *testing.T) {
	n := HexUint64(50)
	hexStr := n.Hex()
	if hexStr != "0x1" {
		t.Errorf("want=%v, got=%v", "0x1", hexStr)
	}
}
