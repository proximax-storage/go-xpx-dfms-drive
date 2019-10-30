package drive

import (
	"github.com/ipfs/go-cid"
	"github.com/libp2p/go-libp2p-core/crypto"
	"github.com/libp2p/go-libp2p-core/peer"

	pb "github.com/proximax-storage/go-xpx-dfms-drive/pb"
)

// UnmarshalID deserializes Id from bytes.
func UnmarshalID(data []byte) (ID, error) {
	id, err := cid.Cast(data)
	if err != nil {
		return cid.Undef, err
	}

	return id, nil
}

// MarshalInvite serializes Invite to bytes using protobuf.
func MarshalInvite(invite Invite) ([]byte, error) {
	return (&pb.Invite{
		Drive:    invite.Drive.Bytes(),
		Owner:    []byte(invite.Owner),
		Duration: invite.Duration,
		Space:    invite.Space,
		Created:  invite.Created,
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
		Created:  proto.Created,
		Duration: proto.Duration,
		Space:    proto.Space,
	}

	invite.Drive, err = UnmarshalID(proto.Drive)
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
	}

	contract.Drive, err = UnmarshalID(proto.Drive)
	if err != nil {
		return
	}

	contract.Root, err = cid.Cast(proto.Root)
	if err != nil {
		return
	}

	contract.Owner, err = peer.IDFromBytes(proto.Owner)
	if err != nil {
		return
	}

	contract.ContractId, err = crypto.UnmarshalPublicKey(proto.ContractId)
	if err != nil {
		return
	}

	for i, m := range proto.Members {
		contract.Members[i], err = peer.IDFromBytes(m)
		if err != nil {
			return
		}
	}
	return
}

func MarshalContract(contract *Contract) ([]byte, error) {
	ctr, err := contract.PrivateKey.GetPublic().Bytes()
	if err != nil {
		return nil, err
	}

	proto := &pb.Contract{
		Drive:            contract.Drive.Bytes(),
		Owner:            []byte(contract.Owner),
		Duration:         contract.Duration,
		Space:            contract.Space,
		Root:             contract.Root.Bytes(),
		ContractId:       ctr,
		Replicas:         uint32(contract.Replicas),
		MinReplicators:   uint32(contract.MinReplicators),
		PercentApprovers: uint32(contract.PercentApprovers),
		Members:          make([][]byte, len(contract.Members)),
	}

	for i, m := range contract.Members {
		proto.Members[i] = []byte(m)
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
	Drive            ID        `json:"drive"`
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
