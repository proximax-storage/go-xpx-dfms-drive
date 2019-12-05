package drive

import (
	"encoding/hex"
	"encoding/json"

	"github.com/ipfs/go-cid"
	"github.com/libp2p/go-libp2p-core/crypto"

	pb "github.com/proximax-storage/go-xpx-dfms-drive/pb"
)

func MarshalContractProto(in *Contract) (_ []byte, err error) {
	out := &pb.Contract{
		Drive:            in.Drive.Bytes(),
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

	out.Owner, err = crypto.MarshalPublicKey(in.Owner)
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

func UnmarshalContractProto(data []byte) (*Contract, error) {
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

	out.Owner, err = crypto.UnmarshalPublicKey(in.Owner)
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
		Drive:            in.Drive.String(),
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

	owner, err := crypto.MarshalPublicKey(in.Owner)
	if err != nil {
		return nil, err
	}
	out.Owner = hex.EncodeToString(owner)

	for i, r := range in.Replicators {
		repl, err := crypto.MarshalPublicKey(r)
		if err != nil {
			return nil, err
		}
		out.Replicators[i] = hex.EncodeToString(repl)
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

	owner, err := hex.DecodeString(in.Owner)
	if err != nil {
		return nil, err
	}

	out.Owner, err = crypto.UnmarshalPublicKey(owner)
	if err != nil {
		return nil, err
	}

	for i, r := range in.Replicators {
		owner, err := hex.DecodeString(r)
		if err != nil {
			return nil, err
		}

		out.Replicators[i], err = crypto.UnmarshalPublicKey(owner)
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

type contractJSON struct {
	Drive            string   `json:"drive"`
	Owner            string   `json:"owner"`
	Replicators      []string `json:"replicators"`
	Root             string   `json:"root"`
	Created          int64    `json:"created"`
	Duration         int64    `json:"duration"`
	Space            int64    `json:"space"`
	PayedReplicas    uint16   `json:"replicas"`
	MinReplicators   uint16   `json:"minReplicators"`
	PercentApprovers uint8    `json:"percentApprovers"`
	BillingPrice     int64    `json:"billingPrice"`
	BillingPeriod    int64    `json:"billingPeriod"`
}
