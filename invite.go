package drive

import (
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
