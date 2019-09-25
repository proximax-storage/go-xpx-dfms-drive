package drive

import (
	"testing"

	"github.com/ipfs/go-cid"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/test"
)

func RandID(t *testing.T) ID {
	id, err := cid.Decode("f015512209d8453505bdc6f269678e16b3e56c2a2948a41f2c792617cc9611ed363c95b63")
	if err != nil {
		t.Fatal(err)
	}

	return id
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
