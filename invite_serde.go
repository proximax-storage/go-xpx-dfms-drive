package drive

import (
	"encoding/hex"
	"encoding/json"

	"github.com/libp2p/go-libp2p-core/crypto"

	pb "github.com/proximax-storage/go-xpx-dfms-drive/pb"
)

// MarshalInviteProto serializes Invite to bytes using protobuf.
func MarshalInviteProto(in *Invite) (_ []byte, err error) {
	out := &pb.Invite{
		Drive:            in.Drive.Bytes(),
		Duration:         in.Duration,
		Space:            in.Space,
		PayedReplicas:    uint32(in.PayedReplicas),
		MinReplicators:   uint32(in.MinReplicators),
		PercentApprovers: uint32(in.PercentApprovers),
		BillingPeriod:    in.BillingPeriod,
		BillingPrice:     in.BillingPrice,
	}

	out.Owner, err = crypto.MarshalPublicKey(in.Owner)
	if err != nil {
		return nil, err
	}

	return out.Marshal()
}

// UnmarshalInviteProto deserializes Invite from bytes using protobuf.
func UnmarshalInviteProto(data []byte) (*Invite, error) {
	in := new(pb.Invite)
	err := in.Unmarshal(data)
	if err != nil {
		return nil, err
	}

	out := &Invite{
		Duration:         in.Duration,
		Space:            in.Space,
		PayedReplicas:    uint16(in.PayedReplicas),
		MinReplicators:   uint16(in.MinReplicators),
		PercentApprovers: uint8(in.PercentApprovers),
		BillingPeriod:    in.BillingPeriod,
		BillingPrice:     in.BillingPrice,
	}

	out.Drive, err = IDFromBytes(in.Drive)
	if err != nil {
		return nil, err
	}

	out.Owner, err = crypto.UnmarshalPublicKey(in.Owner)
	if err != nil {
		return nil, err
	}

	return out, nil
}

func MarshalInviteJSON(in *Invite) (_ []byte, err error) {
	out := &inviteJSON{
		Drive:            in.Drive.String(),
		Duration:         in.Duration,
		Space:            in.Space,
		PayedReplicas:    in.PayedReplicas,
		MinReplicators:   in.MinReplicators,
		PercentApprovers: in.PercentApprovers,
		BillingPrice:     in.BillingPrice,
		BillingPeriod:    in.BillingPeriod,
	}

	owner, err := crypto.MarshalPublicKey(in.Owner)
	if err != nil {
		return nil, err
	}
	out.Owner = hex.EncodeToString(owner)

	return json.Marshal(out)
}

func UnmarshalInviteJSON(data []byte) (*Invite, error) {
	in := &inviteJSON{}
	err := json.Unmarshal(data, in)
	if err != nil {
		return nil, err
	}

	out := &Invite{
		Duration:         in.Duration,
		Space:            in.Space,
		BillingPeriod:    in.BillingPeriod,
		BillingPrice:     in.BillingPrice,
		PayedReplicas:    in.PayedReplicas,
		MinReplicators:   in.MinReplicators,
		PercentApprovers: in.PercentApprovers,
	}

	out.Drive, err = IDFromString(in.Drive)
	if err != nil {
		return nil, err
	}

	owner, err := hex.DecodeString(in.Owner)
	if err != nil {
		return nil, err
	}

	out.Owner, err = crypto.UnmarshalPublicKey(owner)
	if err != nil {
		return nil, err
	}

	return out, nil
}

type inviteJSON struct {
	Drive            string `json:"drive"`
	Owner            string `json:"owner"`
	Duration         int64  `json:"duration"`
	Space            int64  `json:"space"`
	PayedReplicas    uint16 `json:"payedReplicas"`
	MinReplicators   uint16 `json:"minReplicators"`
	PercentApprovers uint8  `json:"percentApprovers"`
	BillingPrice     int64  `json:"billingPrice"`
	BillingPeriod    int64  `json:"billingPeriod"`
}
