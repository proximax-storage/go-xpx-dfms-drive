package drive

import (
	"crypto/rand"
	"io"
	"io/ioutil"
	"testing"

	"github.com/ipfs/go-cid"
	"github.com/libp2p/go-libp2p-core/crypto"
	"github.com/multiformats/go-multihash"
)

func RandID(t *testing.T) ID {
	_, id, err := crypto.GenerateEd25519Key(rand.Reader)
	if err != nil {
		t.Fatal(err)
	}

	return id
}

func RandInvite(t *testing.T) *Invite {
	return &Invite{
		Drive:            RandID(t),
		Owner:            RandID(t),
		Duration:         100,
		Space:            100,
		BillingPrice:     100,
		BillingPeriod:    100,
		PayedReplicas:    100,
		MinReplicators:   100,
		PercentApprovers: 100,
	}
}

func RandContract(t *testing.T) *Contract {
	return &Contract{
		Drive: RandID(t),
		Root:  randCID(t),
		Owner: RandID(t),
		Replicators: []crypto.PubKey{
			RandID(t),
			RandID(t),
		},
		Created:          100,
		Duration:         100,
		Space:            100,
		BillingPrice:     100,
		BillingPeriod:    100,
		PayedReplicas:    100,
		MinReplicators:   100,
		PercentApprovers: 100,
	}
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

	return cid.NewCidV1(cid.Raw, hash)
}
