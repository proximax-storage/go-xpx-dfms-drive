package drive

import (
	"encoding/hex"
	"testing"

	"github.com/libp2p/go-libp2p-core/crypto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestIDFromToPubKey(t *testing.T) {
	in := randKey(t)

	id, err := IDFromPubKey(in)
	require.Nil(t, err, err)

	out, err := IDToPubKey(id)
	require.Nil(t, err, err)

	assert.Equal(t, in, out)
}

func TestIDMarshalUnmarshalBinary(t *testing.T) {
	in := RandID(t)

	data, err := in.MarshalBinary()
	require.Nil(t, err, err)

	out := ID{}
	err = out.UnmarshalBinary(data)
	require.Nil(t, err, err)

	assert.Equal(t, in, out)
}

func TestIDMarshalUnmarshalJSON(t *testing.T) {
	in := RandID(t)

	data, err := in.MarshalJSON()
	require.Nil(t, err, err)

	out := ID{}
	err = out.UnmarshalJSON(data)
	require.Nil(t, err, err)

	assert.Equal(t, in, out)
}

func TestIDFromString(t *testing.T) {
	t.Run("CID", func(t *testing.T) {
		in := "baegbeibondkkrhxfprzwrlgxxltavqhweh2ylhu4hgo5lxjxpqbpfsw2lu"

		out, err := IDFromString(in)
		require.Nil(t, err, err)
		assert.Equal(t, in, out.String())
	})

	t.Run("PublicKey", func(t *testing.T) {
		keyStr := "08011220201b155bf3ebe4dcca522549a9835a21d010d07e6c354f0df30a0a0504b83f1b"

		in, err := IDFromString(keyStr)
		require.Nil(t, err, err)

		keyBts, err := hex.DecodeString(keyStr)
		require.Nil(t, err, err)

		key, err := crypto.UnmarshalPublicKey(keyBts)
		require.Nil(t, err, err)

		out, err := IDFromPubKey(key)
		require.Nil(t, err, err)

		assert.True(t, in.Equals(out))
	})
}
