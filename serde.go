package drive

import (
	"io"

	"github.com/Wondertan/go-serde"
	"github.com/ipfs/go-cid"
	"github.com/libp2p/go-libp2p-core/crypto"
	"github.com/libp2p/go-libp2p-core/peer"

	pb "github.com/proximax-storage/go-xpx-dfms-drive/pb"
)

// MarshalID serializes drive.ID to bytes.
func MarshalID(id ID) ([]byte, error) {
	return id.MarshalBinary()
}

// UnmarshalID deserializes Id from bytes.
func UnmarshalID(data []byte) (ID, error) {
	id, err := cid.Cast(data)
	if err != nil {
		return cid.Undef, err
	}

	return id, nil
}

// MarshalBasicContract serializes BasicContract to bytes using protobuf.
func MarshalBasicContract(basic *BasicContract) ([]byte, error) {
	ctr, err := basic.contractId.Bytes()
	if err != nil {
		return nil, err
	}

	proto := &pb.Contract{
		Drive:      basic.drive.Bytes(),
		Owner:      []byte(basic.owner),
		Members:    make([][]byte, len(basic.members)),
		Duration:   basic.duration,
		Created:    basic.created,
		Space:      basic.space,
		Root:       basic.root.Bytes(),
		ContractId: ctr,
	}

	for i, m := range basic.members {
		proto.Members[i] = []byte(m)
	}

	return proto.Marshal()
}

// UnmarshalBasicContract deserializes BasicContract from bytes using protobuf.
func UnmarshalBasicContract(data []byte) (*BasicContract, error) {
	proto := new(pb.Contract)
	err := proto.Unmarshal(data)
	if err != nil {
		return nil, err
	}

	return protoToBasicContract(proto)
}

// WriteBasicContract serializes invite to the Writer
func WriteBasicContract(w io.Writer, basic *BasicContract) error {
	ctr, err := basic.contractId.Bytes()
	if err != nil {
		return err
	}
	proto := &pb.Contract{
		Drive:      basic.drive.Bytes(),
		Owner:      []byte(basic.owner),
		Members:    make([][]byte, len(basic.members)),
		Duration:   basic.duration,
		Created:    basic.created,
		Space:      basic.space,
		Root:       basic.root.Bytes(),
		ContractId: ctr,
	}

	for i, m := range basic.members {
		proto.Members[i] = []byte(m)
	}

	return serde.WriteMessage(w, proto)
}

// ReadBasicContract deserialize invite from Reader
func ReadBasicContract(r io.Reader) (*BasicContract, error) {
	proto := new(pb.Contract)
	err := serde.ReadMessage(r, proto)
	if err != nil {
		return nil, err
	}

	return protoToBasicContract(proto)
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

// WriteInvite serializes invite to the Writer
func WriteInvite(w io.Writer, invite Invite) error {
	return serde.WriteMessage(w, &pb.Invite{
		Drive:    invite.Drive.Bytes(),
		Owner:    []byte(invite.Owner),
		Duration: invite.Duration,
		Space:    invite.Space,
		Created:  invite.Created,
	})
}

// ReadInvite deserialize invite from Reader
func ReadInvite(r io.Reader) (Invite, error) {
	proto := new(pb.Invite)
	err := serde.ReadMessage(r, proto)
	if err != nil {
		return NilInvite, err
	}

	return protoToInvite(proto)
}

func protoToBasicContract(proto *pb.Contract) (basic *BasicContract, err error) {
	basic = &BasicContract{
		duration: proto.Duration,
		created:  proto.Created,
		space:    proto.Space,
		members:  make([]peer.ID, len(proto.Members)),
	}

	basic.drive, err = UnmarshalID(proto.Drive)
	if err != nil {
		return
	}

	basic.root, err = cid.Cast(proto.Root)
	if err != nil {
		return
	}

	basic.owner, err = peer.IDFromBytes(proto.Owner)
	if err != nil {
		return
	}

	basic.contractId, err = crypto.UnmarshalPublicKey(proto.ContractId)
	if err != nil {
		return
	}

	for i, m := range proto.Members {
		basic.members[i], err = peer.IDFromBytes(m)
		if err != nil {
			return
		}
	}
	return
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

type basicContractJSON struct {
	Drive      ID            `json:"drive"`
	Owner      peer.ID       `json:"owner"`
	Members    []peer.ID     `json:"members"`
	Duration   uint64        `json:"duration"`
	Created    uint64        `json:"created"`
	Root       cid.Cid       `json:"root"`
	Space      uint64        `json:"space"`
	ContractId crypto.PubKey `json:"contractId"`
}
