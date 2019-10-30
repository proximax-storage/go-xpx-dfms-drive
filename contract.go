package drive

import (
	"encoding/json"

	"github.com/ipfs/go-cid"
	"github.com/libp2p/go-libp2p-core/crypto"
	"github.com/libp2p/go-libp2p-core/peer"
)

type PaymentInformation struct {
	Receiver *crypto.PubKey
	Amount   uint64
	Height   uint64
}

type ReplicatorInfo struct {
	Start               uint64
	End                 uint64
	Deposit             uint64
	FilesWithoutDeposit map[cid.Cid]uint16
}

type BillingDescription struct {
	Start    uint64
	End      uint64
	Payments []*PaymentInformation
}

type FileInfo struct {
	FileSize uint64
	Deposit  uint64
	Payments []*PaymentInformation
}

// Contract is an agreement between client and replicator peers on some amount of disk Space
type Contract struct {
	State uint8

	// Drive returns identifier of Contract's Drive
	Drive ID

	// Owner is a peer which owns the Drive
	Owner peer.ID

	// Root is a Drive of Drive's directory
	// Formerly, represents the identifier of all content stored on Drive.
	Root cid.Cid

	// Duration of the contract.
	// NOTE: May return time in milliseconds or amount of blocks,
	// depending on the implementation.
	Duration      int64
	BillingPeriod int64
	BillingPrice  int64

	// Space specifies total physical Space used by Drive
	// on member nodes.
	Space int64

	Replicas uint16

	MinReplicators uint16

	PercentApprovers uint8

	Replicators map[crypto.PubKey]*ReplicatorInfo

	BillingHistory []BillingDescription

	Files map[cid.Cid]*FileInfo

	ContractId crypto.PubKey

	// PrivateKey specifies private key for signing
	PrivateKey crypto.PrivKey

	// Created is a Drive starting moment.
	// Usually derived from ab invitation
	// NOTE: May return time in milliseconds or amount of blocks,
	// depending on the implementation.
	Created int64

	// Members are peers responsible for Drive
	// Members() []peer.ID
	Members []peer.ID
}

func NewContractFromInvite(
	invite Invite,
	root cid.Cid,
	contractId crypto.PubKey,
	privateKey crypto.PrivKey,
	members ...peer.ID,
) *Contract {
	return &Contract{
		Drive:            invite.Drive,
		Owner:            invite.Owner,
		Duration:         invite.Duration,
		Root:             root,
		Space:            invite.Space,
		Replicas:         invite.Replicas,
		MinReplicators:   invite.MinReplicators,
		PercentApprovers: invite.PercentApprovers,
		BillingPeriod:    invite.BillingPeriod,
		BillingPrice:     invite.BillingPrice,

		ContractId: contractId,
		PrivateKey: privateKey,
		Members:    members,
	}
}

func NewContract(
	state uint8,
	drive ID,
	owner peer.ID,
	duration int64,
	billingPeriod int64,
	billingPrice int64,
	space int64,
	replicas uint16,
	minReplicators uint16,
	percentApprovers uint8,
	replicators map[crypto.PubKey]*ReplicatorInfo,
	contractId crypto.PubKey,
	members ...peer.ID,

) *Contract {
	return &Contract{
		State:            state,
		Drive:            drive,
		Owner:            owner,
		Duration:         duration,
		Space:            space,
		Replicas:         replicas,
		MinReplicators:   minReplicators,
		PercentApprovers: percentApprovers,
		BillingPeriod:    billingPeriod,
		BillingPrice:     billingPrice,
		Replicators:      replicators,
		Members:          members,
		ContractId:       contractId,
	}
}

func NewEmptyContract() *Contract {
	return new(Contract)
}

func (c *Contract) MarshalBinary() ([]byte, error) {
	return MarshalContract(c)
}

func (c *Contract) UnmarshalBinary(data []byte) error {
	basic, err := UnmarshalContract(data)
	if err != nil {
		return err
	}

	*c = *basic
	return nil
}

func (c *Contract) MarshalJSON() ([]byte, error) {
	contractId, err := c.ContractId.Bytes()
	if err != nil {
		return nil, err
	}

	return json.Marshal(&contractJSON{
		Drive:            c.Drive,
		Owner:            c.Owner,
		Duration:         c.Duration,
		Root:             c.Root,
		Space:            c.Space,
		ContractId:       contractId,
		Replicas:         c.Replicas,
		MinReplicators:   c.MinReplicators,
		PercentApprovers: c.PercentApprovers,
		Members:          c.Members,
	})
}

func (c *Contract) UnmarshalJSON(data []byte) error {
	cjson := &contractJSON{}
	err := json.Unmarshal(data, cjson)
	if err != nil {
		return err
	}

	contractId, err := crypto.UnmarshalPublicKey(cjson.ContractId)
	if err != nil {
		return err
	}
	c.Drive = cjson.Drive
	c.Owner = cjson.Owner
	c.Duration = cjson.Duration
	c.Root = cjson.Root
	c.Space = cjson.Space
	c.ContractId = contractId
	c.Replicas = cjson.Replicas
	c.MinReplicators = cjson.MinReplicators
	c.PercentApprovers = cjson.PercentApprovers
	c.Members = cjson.Members

	return nil
}
