package drive

import (
	"encoding/json"
	"github.com/ipfs/go-cid"
	"github.com/libp2p/go-libp2p-core/crypto"
	"github.com/libp2p/go-libp2p-core/peer"

	pb "github.com/proximax-storage/go-xpx-dfms-drive/pb"
)

// MarshalInvite serializes Invite to bytes using protobuf.
func MarshalInvite(invite Invite) ([]byte, error) {
	id, err := IdToBytes(invite.Drive)
	if err != nil {
		return nil, err
	}

	return (&pb.Invite{
		Drive:            id,
		Owner:            []byte(invite.Owner),
		Duration:         invite.Duration,
		Space:            invite.Space,
		Created:          invite.Created,
		Replicas:         uint32(invite.Replicas),
		MinReplicators:   uint32(invite.MinReplicators),
		PercentApprovers: uint32(invite.PercentApprovers),
	}).Marshal()
}

// UnmarshalInvite deserializes Invite from bytes using protobuf.
func UnmarshalInvite(data []byte) (Invite, error) {
	proto := new(pb.Invite)
	err := proto.Unmarshal(data)
	if err != nil {
		return NilInvite, err
	}

	return protoToInvite(proto)
}

func protoToInvite(proto *pb.Invite) (invite Invite, err error) {
	invite = Invite{
		Created:          proto.Created,
		Duration:         proto.Duration,
		Space:            proto.Space,
		Replicas:         uint16(proto.Replicas),
		MinReplicators:   uint16(proto.MinReplicators),
		PercentApprovers: uint8(proto.PercentApprovers),
	}

	invite.Drive, err = IdFromBytes(proto.Drive)
	if err != nil {
		return NilInvite, err
	}

	invite.Owner, err = peer.IDFromBytes(proto.Owner)
	if err != nil {
		return NilInvite, err
	}

	return
}

func protoToContract(proto *pb.Contract) (contract *Contract, err error) {
	contract = &Contract{
		Duration:         proto.Duration,
		Space:            proto.Space,
		Replicas:         uint16(proto.Replicas),
		MinReplicators:   uint16(proto.MinReplicators),
		PercentApprovers: uint8(proto.PercentApprovers),
		BillingPrice:     proto.BillingPrice,
		BillingPeriod:    proto.BillingPeriod,
		Members:          make([]peer.ID, len(proto.Members)),
		Replicators:      make(map[crypto.PubKey]*ReplicatorInfo, len(proto.Replicators)),
	}

	contract.Drive, err = IdFromBytes(proto.Drive)
	if err != nil {
		return nil, err
	}

	contract.Root, err = cid.Cast(proto.Root)
	if err != nil {
		return nil, err
	}

	contract.Owner, err = peer.IDFromBytes(proto.Owner)
	if err != nil {
		return nil, err
	}

	for i, m := range proto.Members {
		contract.Members[i], err = peer.IDFromBytes(m)
		if err != nil {
			return nil, err
		}
	}

	for _, r := range proto.Replicators {
		pubKey, err := IdFromBytes(r.PubKey)
		if err != nil {
			return nil, err
		}
		var rInfo *ReplicatorInfo
		err = json.Unmarshal(r.Info, &rInfo)
		if err != nil {
			return nil, err
		}
		contract.Replicators[pubKey] = rInfo
	}
	return
}

func MarshalContract(contract *Contract) ([]byte, error) {

	id, err := IdToBytes(contract.Drive)
	if err != nil {
		return nil, err
	}

	proto := &pb.Contract{
		Drive:            id,
		Owner:            []byte(contract.Owner),
		Duration:         contract.Duration,
		Space:            contract.Space,
		Root:             contract.Root.Bytes(),
		Replicas:         uint32(contract.Replicas),
		MinReplicators:   uint32(contract.MinReplicators),
		PercentApprovers: uint32(contract.PercentApprovers),
		Members:          make([][]byte, len(contract.Members)),
		Replicators:      make([]*pb.ReplicatorInfo, len(contract.Replicators)),
	}

	for i, m := range contract.Members {
		proto.Members[i] = []byte(m)
	}

	iter := 0
	for pk, r := range contract.Replicators {
		pubKey, err := IdToBytes(pk)
		if err != nil {
			return nil, err
		}
		// TODO: move to Protobuf
		rInfo, err := json.Marshal(r)
		if err != nil {
			return nil, err
		}
		proto.Replicators[iter] = &pb.ReplicatorInfo{
			PubKey: pubKey,
			Info:   rInfo,
		}
		iter++
	}

	return proto.Marshal()
}

func UnmarshalContract(data []byte) (*Contract, error) {
	proto := new(pb.Contract)
	err := proto.Unmarshal(data)
	if err != nil {
		return nil, err
	}

	return protoToContract(proto)
}

type contractJSON struct {
	Drive            []byte    `json:"drive"`
	Owner            peer.ID   `json:"owner"`
	Duration         int64     `json:"duration"`
	Root             cid.Cid   `json:"root"`
	Space            int64     `json:"space"`
	ContractId       []byte    `json:"contractId"`
	Replicas         uint16    `json:"replicas"`
	MinReplicators   uint16    `json:"minReplicators"`
	PercentApprovers uint8     `json:"percentApprovers"`
	Members          []peer.ID `json:"members"`
}
