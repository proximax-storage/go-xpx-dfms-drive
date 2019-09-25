package drive

import (
	"github.com/ipfs/go-cid"
	"github.com/libp2p/go-libp2p-core/peer"
)

var NilInvite = Invite{}

// Invite represents invitation for replicators to join Contract as a member.
type Invite struct {
	Drive    ID
	Created  uint64
	Duration uint64
	Space    uint64
	Owner    peer.ID
}

// ID returns Drive ID of the Invite and makes new if none.
func (i *Invite) ID() (ID, error) {
	if i.Drive != cid.Undef {
		return i.Drive, nil
	}

	data, err := MarshalInvite(*i)
	if err != nil {
		return cid.Undef, err
	}

	id, err := NewIDFromBytes(data)
	if err != nil {
		return cid.Undef, err
	}

	i.Drive = id
	return id, nil
}
