package drive

import (
	"encoding/json"
	"github.com/ipfs/go-cid"
	"github.com/libp2p/go-libp2p-core/crypto"
	"github.com/libp2p/go-libp2p-core/peer"
)

type LedgerContract struct {
	drive               ID
	owner               peer.ID
	duration            uint64
	created             uint64
	root                cid.Cid
	space               uint64
	contractId          crypto.PubKey
	privateKey          crypto.PrivKey
	replicasDelta       int8
	minReplicatorsDelta int8
	minApproversDelta   int8
}

func NewLedgerContractFromInvite(
	invite Invite,
	root cid.Cid,
	contractId crypto.PubKey,
	privateKey crypto.PrivKey,
) *LedgerContract {
	return &LedgerContract{
		drive:               invite.Drive,
		owner:               invite.Owner,
		duration:            invite.Duration,
		created:             invite.Created,
		root:                root,
		space:               invite.Space,
		contractId:          contractId,
		privateKey:          privateKey,
		replicasDelta:       invite.ReplicasDelta,
		minReplicatorsDelta: invite.MinReplicatorsDelta,
		minApproversDelta:   invite.MinApproversDelta,
	}
}

func (c *LedgerContract) Drive() ID {
	return c.drive
}

func (c *LedgerContract) Owner() peer.ID {
	return c.owner
}

func (c *LedgerContract) ContractID() crypto.PubKey {
	return c.contractId
}

func (c *LedgerContract) PrivateKey() crypto.PrivKey {
	return c.privateKey
}

func (c *LedgerContract) Duration() uint64 {
	return c.duration
}

func (c *LedgerContract) Created() uint64 {
	return c.created
}

func (c *LedgerContract) Root() cid.Cid {
	return c.root
}

func (c *LedgerContract) TotalSpace() uint64 {
	return c.space
}

func (c *LedgerContract) Replicas() int8 {
	return c.replicasDelta
}

func (c *LedgerContract) MinReplicatorsDelta() int8 {
	return c.minReplicatorsDelta
}

func (c *LedgerContract) MinApproversDelta() int8 {
	return c.minApproversDelta
}

func (c *LedgerContract) MarshalBinary() ([]byte, error) {
	return MarshalLedgerContract(c)
}

func (c *LedgerContract) UnmarshalBinary(data []byte) error {
	basic, err := UnmarshalLedgerContract(data)
	if err != nil {
		return err
	}

	*c = *basic
	return nil
}

func (c *LedgerContract) MarshalJSON() ([]byte, error) {
	contractId, err := c.contractId.Bytes()
	if err != nil {
		return nil, err
	}

	return json.Marshal(&ledgerContractJSON{
		Drive:               c.drive,
		Owner:               c.owner,
		Duration:            c.duration,
		Created:             c.created,
		Root:                c.root,
		Space:               c.space,
		ContractId:          contractId,
		Replicas:            c.replicasDelta,
		MinReplicatorsDelta: c.minReplicatorsDelta,
		MinApproversDelta:   c.minApproversDelta,
	})
}

func (c *LedgerContract) UnmarshalJSON(data []byte) error {
	cjson := &ledgerContractJSON{}
	err := json.Unmarshal(data, cjson)
	if err != nil {
		return err
	}

	contractId, err := crypto.UnmarshalPublicKey(cjson.ContractId)
	if err != nil {
		return err
	}
	c.drive = cjson.Drive
	c.owner = cjson.Owner
	c.duration = cjson.Duration
	c.created = cjson.Created
	c.root = cjson.Root
	c.space = cjson.Space
	c.contractId = contractId
	c.replicasDelta = cjson.Replicas
	c.minReplicatorsDelta = cjson.MinReplicatorsDelta
	c.minApproversDelta = cjson.MinApproversDelta

	return nil
}
