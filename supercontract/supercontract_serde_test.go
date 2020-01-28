package supercontract

import (
	"crypto/rand"
	"github.com/ipfs/go-cid"
	"github.com/multiformats/go-multihash"
	drive "github.com/proximax-storage/go-xpx-dfms-drive"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"io/ioutil"
	"testing"
)

func TestMarshalUnmarshalBinarySuperContract(t *testing.T) {
	in := randSuperContract(t)

	data, err := in.MarshalBinary()
	require.Nil(t, err, err)

	out := &SuperContract{}
	err = out.UnmarshalBinary(data)
	require.Nil(t, err, err)

	assert.Equal(t, in, out)
}

func TestMarshalUnmarshalJSONSuperContract(t *testing.T) {
	in := randSuperContract(t)

	data, err := in.MarshalJSON()
	require.Nil(t, err, err)

	out := &SuperContract{}
	err = out.UnmarshalJSON(data)
	require.Nil(t, err, err)

	assert.Equal(t, in, out)
}

func randSuperContract(t *testing.T) *SuperContract {
	return &SuperContract{
		ID:        randID(t),
		Drive:     drive.RandContract(t),
		File:      randCID(t),
		VMVersion: 3,
		Functions: []string{"main"},
	}
}

func randID(t *testing.T) ID {
	return ID(randCID(t))
}

func randCID(t *testing.T) cid.Cid {
	b, err := ioutil.ReadAll(io.LimitReader(rand.Reader, 256))
	if err != nil {
		t.Fatal(err)
	}

	hash, err := multihash.Sum(b, multihash.SHA2_256, -1)
	if err != nil {
		t.Fatal(err)
	}

	return cid.NewCidV1(codec, hash)
}
