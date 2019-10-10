package drive

import (
	"crypto/rand"
	"io"
	"io/ioutil"
	"testing"

	"github.com/ipfs/go-cid"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/test"
	"github.com/multiformats/go-multihash"
)

func RandID(t *testing.T) ID {
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

func RandInvite(t *testing.T) Invite {
	return Invite{
		Drive:    RandID(t),
		Created:  100,
		Duration: 100,
		Space:    100,
		Owner:    test.RandPeerIDFatal(t),
	}
}

func RandBasicContract(t *testing.T) *BasicContract {
	id := RandID(t)
	return &BasicContract{
		drive: id,
		root:  id,
		owner: test.RandPeerIDFatal(t),
		members: []peer.ID{
			test.RandPeerIDFatal(t),
			test.RandPeerIDFatal(t),
		},
		duration: 100,
		created:  100,
		space:    100,
	}
}
