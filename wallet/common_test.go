package wallet

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVariables(t *testing.T) {
	assert.Equal(t, "0x0000000000000000", daoDepositOutputDataHex.Hex())
}
