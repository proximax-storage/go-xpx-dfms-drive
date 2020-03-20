package supercontract

import (
	"crypto/rand"
	"io"
	"io/ioutil"
	"testing"

	"github.com/ipfs/go-cid"
	"github.com/multiformats/go-multihash"
)

func RandID(t *testing.T) ID {
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
