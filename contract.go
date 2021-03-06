package drive

import (
	"github.com/ipfs/go-cid"
	"github.com/libp2p/go-libp2p-core/crypto"
)

// TODO Add payment/billing information as separate structure

// Role defines behavior types of members in Drive Contracts.
type Role uint8

const (
	Owner Role = iota
	Replicator
)

// Billet is a preliminary Contract or simply uninitialized Contract
type Billet struct {
	Invite     Invite
	PrivateKey crypto.PrivKey
}

// UploadInfo for contract files
type UploadInfo struct {
	Participant  crypto.PubKey
	UploadedSize int
}

// TODO Docs
// TODO Still needs reconsideration about some fields.
// Contract is an agreement between client and replicator peers on some amount of disk Space
type Contract struct {
	// Drive returns identifier of Contract's Drive
	Drive ID

	// Owner of the Contract. Has write access on the Contract's Drive and can control it's state.
	Owner crypto.PubKey

	// Replicators are nodes which responsible for replication of DriveFS. Represented as public keys.
	Replicators []crypto.PubKey

	// Root is a CID of Drive's top-level directory.
	// It is used as an entry point to access DriveFS.
	Root cid.Cid

	// Block height when the Contract was started.
	Created int64

	// Duration of the contract.
	Duration      int64
	BillingPeriod int64
	BillingPrice  int64

	// Space specifies total physical Space used by Drive
	// on replicator nodes.
	Space int64

	// TODO Rename all of this
	PayedReplicas    uint16
	MinReplicators   uint16
	PercentApprovers uint8

	PrivateKey crypto.PrivKey // TODO Use Billet instead
}

func (c *Contract) MarshalBinary() ([]byte, error) {
	return MarshalContractProto(c)
}

func (c *Contract) UnmarshalBinary(data []byte) error {
	out, err := UnmarshalContractProto(data)
	if err != nil {
		return err
	}

	*c = *out
	return nil
}

func (c *Contract) MarshalJSON() ([]byte, error) {
	return MarshalContractJSON(c)
}

func (c *Contract) UnmarshalJSON(data []byte) error {
	out, err := UnmarshalContractJSON(data)
	if err != nil {
		return err
	}

	*c = *out
	return nil
}
