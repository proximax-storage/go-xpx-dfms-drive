package drive

import (
	"crypto/rand"
	"io"
	"io/ioutil"
	"testing"

	"github.com/ipfs/go-cid"
	"github.com/libp2p/go-libp2p-core/crypto"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/test"
	"github.com/multiformats/go-multihash"
)

func RandCID(t *testing.T) cid.Cid {
	b, err := ioutil.ReadAll(io.LimitReader(rand.Reader, 256))
	if err != nil {
		t.Fatal(err)
	}

	hash, err := multihash.Sum(b, multihash.SHA2_256, -1)
	if err != nil {
		t.Fatal(err)
	}

	return cid.NewCidV1(cid.Raw, hash)
}

func RandID(t *testing.T) ID {
	b, err := ioutil.ReadAll(io.LimitReader(rand.Reader, 32))
	if err != nil {
		t.Fatal(err)
	}

	hash, err := NewIDFromBytes(b)
	if err != nil {
		t.Fatal(err)
	}

	return hash
}
func RandInvite(t *testing.T) Invite {
	return Invite{
		Drive:    RandID(t),
		Created:  100,
		Duration: 100,
		Space:    100,
		Owner:    test.RandPeerIDFatal(t),
	}
}

func RandContract(t *testing.T) *Contract {

	privKey, _, err := crypto.GenerateKeyPair(crypto.Ed25519, 1024)
	if err != nil {
		t.Fatal(err)
	}
	return &Contract{
		Drive: RandID(t),
		Root:  RandCID(t),
		Owner: test.RandPeerIDFatal(t),
		Members: []peer.ID{
			test.RandPeerIDFatal(t),
			test.RandPeerIDFatal(t),
		},
		Duration:   100,
		Space:      100,
		PrivateKey: privKey,
	}
}
