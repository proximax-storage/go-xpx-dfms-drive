package drive

import (
	"encoding/hex"
	"encoding/json"

	"github.com/ipfs/go-cid"
	"github.com/libp2p/go-libp2p-core/crypto"

	pb "github.com/proximax-storage/go-xpx-dfms-drive/pb"
)

func IDFromString(data string) (ID, error) {
	dec, err := hex.DecodeString(data)
	if err != nil {
		return nil, err
	}

	return crypto.UnmarshalPublicKey(dec)
}

// TODO Remove key type marshalling
func IDToString(id ID) (string, error) {
	s, err := crypto.MarshalPublicKey(id)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(s), nil
}

func IDFromBytes(data []byte) (ID, error) {
	return crypto.UnmarshalPublicKey(data)
}

func IDToBytes(id ID) ([]byte, error) {
	return crypto.MarshalPublicKey(id)
}

// MarshalInvite serializes Invite to bytes using protobuf.
func MarshalInvite(in *Invite) (_ []byte, err error) {
	out := &pb.Invite{
		Duration:         in.Duration,
		Space:            in.Space,
		PayedReplicas:    uint32(in.PayedReplicas),
		MinReplicators:   uint32(in.MinReplicators),
		PercentApprovers: uint32(in.PercentApprovers),
		BillingPeriod:    in.BillingPeriod,
		BillingPrice:     in.BillingPrice,
	}

	out.Drive, err = IDToBytes(in.Drive)
	if err != nil {
		return nil, err
	}

	out.Owner, err = crypto.MarshalPublicKey(in.Owner)
	if err != nil {
		return nil, err
	}

	return out.Marshal()
}

// UnmarshalInvite deserializes Invite from bytes using protobuf.
func UnmarshalInvite(data []byte) (*Invite, error) {
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

func MarshalContract(in *Contract) (_ []byte, err error) {
	out := &pb.Contract{
		Replicators:      make([][]byte, len(in.Replicators)),
		Root:             in.Root.Bytes(),
		Created:          in.Created,
		BillingPrice:     in.BillingPrice,
		BillingPeriod:    in.BillingPeriod,
		Duration:         in.Duration,
		Space:            in.Space,
		PayedReplicas:    uint32(in.PayedReplicas),
		MinReplicators:   uint32(in.MinReplicators),
		PercentApprovers: uint32(in.PercentApprovers),
	}

	out.Owner, err = IDToBytes(in.Owner)
	if err != nil {
		return nil, err
	}

	out.Drive, err = IDToBytes(in.Drive)
	if err != nil {
		return nil, err
	}

	for i, r := range in.Replicators {
		out.Replicators[i], err = crypto.MarshalPublicKey(r)
		if err != nil {
			return nil, err
		}
	}

	return out.Marshal()
}

func UnmarshalContract(data []byte) (*Contract, error) {
	in := new(pb.Contract)
	err := in.Unmarshal(data)
	if err != nil {
		return nil, err
	}

	out := &Contract{
		Replicators:      make([]crypto.PubKey, len(in.Replicators)),
		Created:          in.Created,
		Duration:         in.Duration,
		Space:            in.Space,
		PayedReplicas:    uint16(in.PayedReplicas),
		MinReplicators:   uint16(in.MinReplicators),
		PercentApprovers: uint8(in.PercentApprovers),
		BillingPrice:     in.BillingPrice,
		BillingPeriod:    in.BillingPeriod,
	}

	out.Drive, err = IDFromBytes(in.Drive)
	if err != nil {
		return nil, err
	}

	out.Root, err = cid.Cast(in.Root)
	if err != nil {
		return nil, err
	}

	out.Owner, err = IDFromBytes(in.Owner)
	if err != nil {
		return nil, err
	}

	for i, r := range in.Replicators {
		out.Replicators[i], err = crypto.UnmarshalPublicKey(r)
		if err != nil {
			return nil, err
		}
	}

	return out, nil
}

func MarshalContractJSON(in *Contract) (_ []byte, err error) {
	out := &contractJSON{
		Replicators:      make([]string, len(in.Replicators)),
		Root:             in.Root.String(),
		Created:          in.Created,
		Duration:         in.Duration,
		Space:            in.Space,
		PayedReplicas:    in.PayedReplicas,
		MinReplicators:   in.MinReplicators,
		PercentApprovers: in.PercentApprovers,
		BillingPeriod:    in.BillingPeriod,
		BillingPrice:     in.BillingPrice,
	}

	out.Drive, err = IDToString(in.Drive)
	if err != nil {
		return nil, err
	}

	out.Owner, err = IDToString(in.Owner)
	if err != nil {
		return nil, err
	}

	for i, r := range in.Replicators {
		out.Replicators[i], err = IDToString(r)
		if err != nil {
			return nil, err
		}
	}

	return json.Marshal(out)
}

func UnmarshalContractJSON(data []byte) (*Contract, error) {
	in := &contractJSON{}
	err := json.Unmarshal(data, in)
	if err != nil {
		return nil, err
	}

	out := &Contract{
		Replicators:      make([]crypto.PubKey, len(in.Replicators)),
		Created:          in.Created,
		Duration:         in.Duration,
		BillingPeriod:    in.BillingPeriod,
		BillingPrice:     in.BillingPrice,
		Space:            in.Space,
		PayedReplicas:    in.PayedReplicas,
		MinReplicators:   in.MinReplicators,
		PercentApprovers: in.PercentApprovers,
	}

	out.Drive, err = IDFromString(in.Drive)
	if err != nil {
		return nil, err
	}

	out.Owner, err = IDFromString(in.Owner)
	if err != nil {
		return nil, err
	}

	for i, r := range in.Replicators {
		out.Replicators[i], err = IDFromString(r)
		if err != nil {
			return nil, err
		}
	}

	out.Root, err = cid.Decode(in.Root)
	if err != nil {
		return nil, err
	}

	return out, nil
}

func MarshalInviteJSON(in *Invite) (_ []byte, err error) {
	out := &inviteJSON{
		Duration:         in.Duration,
		Space:            in.Space,
		PayedReplicas:    in.PayedReplicas,
		MinReplicators:   in.MinReplicators,
		PercentApprovers: in.PercentApprovers,
		BillingPrice:     in.BillingPrice,
		BillingPeriod:    in.BillingPeriod,
	}

	out.Drive, err = IDToString(in.Drive)
	if err != nil {
		return nil, err
	}

	out.Owner, err = IDToString(in.Owner)
	if err != nil {
		return nil, err
	}

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

	out.Owner, err = IDFromString(in.Owner)
	if err != nil {
		return nil, err
	}

	return out, nil
}

type contractJSON struct {
	Drive            string   `json:"drive"`
	Owner            string   `json:"owner"`
	Replicators      []string `json:"replicators"`
	Root             string   `json:"root"`
	Created          int64    `json:"created"`
	Duration         int64    `json:"duration"`
	Space            int64    `json:"space"`
	PayedReplicas    uint16   `json:"payedReplicas"`
	MinReplicators   uint16   `json:"minReplicators"`
	PercentApprovers uint8    `json:"percentApprovers"`
	BillingPrice     int64    `json:"billingPrice"`
	BillingPeriod    int64    `json:"billingPeriod"`
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
