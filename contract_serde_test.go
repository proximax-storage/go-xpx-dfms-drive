package drive

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMarshalUnmarshalBinaryContract(t *testing.T) {
	in := RandContract(t)

	data, err := in.MarshalBinary()
	require.Nil(t, err, err)

	out := &Contract{}
	err = out.UnmarshalBinary(data)
	require.Nil(t, err, err)

	assert.Equal(t, in, out)
}

func TestMarshalUnmarshalJSONContract(t *testing.T) {
	in := RandContract(t)

	data, err := in.MarshalJSON()
	require.Nil(t, err, err)

	out := &Contract{}
	err = out.UnmarshalJSON(data)
	require.Nil(t, err, err)

	assert.Equal(t, in, out)
}
