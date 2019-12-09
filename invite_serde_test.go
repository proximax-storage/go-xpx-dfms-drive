package drive

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMarshalUnmarshalBinaryInvite(t *testing.T) {
	in := RandInvite(t)

	data, err := in.MarshalBinary()
	require.Nil(t, err, err)

	out := &Invite{}
	err = out.UnmarshalBinary(data)
	require.Nil(t, err, err)

	assert.Equal(t, in, out)
}

func TestMarshalUnmarshalJSONInvite(t *testing.T) {
	in := RandInvite(t)

	data, err := in.MarshalJSON()
	require.Nil(t, err, err)

	out := &Invite{}
	err = out.UnmarshalJSON(data)
	require.Nil(t, err, err)

	assert.Equal(t, in, out)
}
