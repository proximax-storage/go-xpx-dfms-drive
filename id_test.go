package drive

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMarshalUnmarshalIdFromString(t *testing.T) {
	in := RandID(t)

	s, err := IdToString(in)
	assert.Nil(t, err)

	out, err := NewIDFromString(s)
	assert.Nil(t, err)
	assert.Equal(t, in, out)
}
