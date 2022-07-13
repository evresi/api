package data

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToFromString(t *testing.T) {
	in := "2dad4587-f251-45b3-bf18-d1f6ba2c421f"

	id := ParseUUID(in)

	assert.True(t, id.Valid, "Created UUID should be valid")

	out := id.String()

	assert.Equal(t, in, out, "Serialized string should be the same as the value that was deserialized")
}
