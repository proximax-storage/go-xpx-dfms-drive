package drive

import (
	"github.com/libp2p/go-libp2p-core/peer"
)

var NilInvite = Invite{}

// Invite represents invitation for Replicators to join Contract as a member.
type Invite struct {
	Drive            ID
	Created          int64
	Duration         int64
	Space            int64
	Owner            peer.ID
	Replicas         uint16
	MinReplicators   uint16
	PercentApprovers uint8
	BillingPeriod    int64
	BillingPrice     int64
}
