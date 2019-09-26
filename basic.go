package drive

import (
	"encoding/json"

	"github.com/ipfs/go-cid"
	"github.com/libp2p/go-libp2p-core/peer"
)

type BasicContract struct {
	drive    ID
	owner    peer.ID
	members  []peer.ID
	duration uint64
	created  uint64
	root     cid.Cid
	space    uint64
}

func NewBasicContractFromInvite(invite Invite, members []peer.ID) *BasicContract {
	return &BasicContract{
		drive:    invite.Drive,
		owner:    invite.Owner,
		members:  members,
		duration: invite.Duration,
		created:  invite.Created,
		root:     cid.Undef,
		space:    invite.Space,
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

	c.drive = cjson.Drive
	c.owner = cjson.Owner
	c.members = cjson.Members
	c.duration = cjson.Duration
	c.created = cjson.Created
	c.root = cjson.Root
	c.space = cjson.Space

	return nil
}
