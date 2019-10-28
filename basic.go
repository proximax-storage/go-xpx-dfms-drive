package drive

import (
	"encoding/json"

	"github.com/ipfs/go-cid"
	"github.com/libp2p/go-libp2p-core/crypto"
	"github.com/libp2p/go-libp2p-core/peer"
)

type BasicContract struct {
	drive               ID
	owner               peer.ID
	members             []peer.ID
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

func NewBasicContractFromInvite(
	invite Invite,
	root cid.Cid,
	contractId crypto.PubKey,
	privateKey crypto.PrivKey,
) *BasicContract {
	return &BasicContract{
		drive:               invite.Drive,
		owner:               invite.Owner,
		duration:            invite.Duration,
		created:             invite.Created,
		root:                root,
		space:               invite.Space,
		privateKey:          privateKey,
		contractId:          contractId,
		replicasDelta:       invite.ReplicasDelta,
		minReplicatorsDelta: invite.MinReplicatorsDelta,
		minApproversDelta:   invite.MinApproversDelta,
	}
}

func NewEmptyBasicContract() *BasicContract {
	return new(BasicContract)
}

func (c *BasicContract) Drive() ID {
	return c.drive
}

func (c *BasicContract) Owner() peer.ID {
	return c.owner
}

func (c *BasicContract) ContractID() crypto.PubKey {
	return c.contractId
}

func (c *BasicContract) PrivateKey() crypto.PrivKey {
	return c.privateKey
}

func (c *BasicContract) Members() []peer.ID {
	return c.members
}

func (c *BasicContract) Duration() uint64 {
	return c.duration
}

func (c *BasicContract) Created() uint64 {
	return c.created
}

func (c *BasicContract) Root() cid.Cid {
	return c.root
}

func (c *BasicContract) TotalSpace() uint64 {
	return c.space
}

func (c *BasicContract) Replicas() int8 {
	return c.replicasDelta
}

func (c *BasicContract) MinReplicatorsDelta() int8 {
	return c.minReplicatorsDelta
}

func (c *BasicContract) MinApproversDelta() int8 {
	return c.minApproversDelta
}

func (c *BasicContract) MarshalBinary() ([]byte, error) {
	return MarshalBasicContract(c)
}

func (c *BasicContract) UnmarshalBinary(data []byte) error {
	basic, err := UnmarshalBasicContract(data)
	if err != nil {
		return err
	}

	*c = *basic
	return nil
}

func (c *BasicContract) MarshalJSON() ([]byte, error) {

	return json.Marshal(&basicContractJSON{
		Drive:    c.drive,
		Owner:    c.owner,
		Members:  c.members,
		Duration: c.duration,
		Created:  c.created,
		Root:     c.root,
		Space:    c.space,
	})
}

func (c *BasicContract) UnmarshalJSON(data []byte) error {
	cjson := &basicContractJSON{}
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
	c.members = cjson.Members
	c.duration = cjson.Duration
	c.created = cjson.Created
	c.root = cjson.Root
	c.space = cjson.Space
	c.replicasDelta = cjson.Replicas
	c.contractId = contractId

	return nil
}
