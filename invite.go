package drive

import (
	"github.com/libp2p/go-libp2p-core/crypto"
)

// Invite represents invitation for Replicators to join Drive Contract.
type Invite struct {
	Drive ID
	Owner crypto.PubKey

	Duration      int64
	Space         int64
	BillingPeriod int64
	BillingPrice  int64

	// TODO Rename
	PayedReplicas    uint16
	MinReplicators   uint16
	PercentApprovers uint8
}

func (in *Invite) MarshalBinary() ([]byte, error) {
	return MarshalInviteProto(in)
}

func (in *Invite) UnmarshalBinary(data []byte) error {
	out, err := UnmarshalInviteProto(data)
	if err != nil {
		return err
	}

	*in = *out
	return nil
}

func (in *Invite) MarshalJSON() ([]byte, error) {
	return MarshalInviteJSON(in)
}

func (in *Invite) UnmarshalJSON(data []byte) error {
	out, err := UnmarshalInviteJSON(data)
	if err != nil {
		return err
	}

	*in = *out
	return nil
}
