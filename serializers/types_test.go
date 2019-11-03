package serializers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTypes_Serialize(t *testing.T) {
	assert.Equal(t, "0x08", Serialize())
}
