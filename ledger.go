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

type LedgerContract struct {
	state            uint8
	drive            ID
	owner            peer.ID
	root             cid.Cid
	duration         int64
	billingPeriod    int64
	billingPrice     int64
	space            int64
	replicas         uint16
	minReplicators   uint16
	percentApprovers uint8

	replicators    map[crypto.PubKey]*ReplicatorInfo
	billingHistory []BillingDescription
	files          map[cid.Cid]*FileInfo

	contractId crypto.PubKey
	privateKey crypto.PrivKey
}

func NewLedgerContractFromInvite(
	invite Invite,
	root cid.Cid,
	contractId crypto.PubKey,
	privateKey crypto.PrivKey,
) *LedgerContract {
	return &LedgerContract{
		drive:            invite.Drive,
		owner:            invite.Owner,
		duration:         invite.Duration,
		root:             root,
		space:            invite.Space,
		replicas:         invite.Replicas,
		minReplicators:   invite.MinReplicators,
		percentApprovers: invite.PercentApprovers,
		billingPeriod:    invite.BillingPeriod,
		billingPrice:     invite.BillingPrice,

		contractId: contractId,
		privateKey: privateKey,
	}
}

func NewLedgerContract(
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
) *LedgerContract {
	return &LedgerContract{
		state:            state,
		drive:            drive,
		owner:            owner,
		duration:         duration,
		space:            space,
		replicas:         replicas,
		minReplicators:   minReplicators,
		percentApprovers: percentApprovers,
		billingPeriod:    billingPeriod,
		billingPrice:     billingPrice,
		replicators:      replicators,

		contractId: contractId,
	}
}

func (c *LedgerContract) State() uint8 {
	return c.state
}

// TODO Need clarify returned data type
func (c *LedgerContract) Replicators() map[crypto.PubKey]*ReplicatorInfo {
	return c.replicators
}

func (c *LedgerContract) BillingPeriod() int64 {
	return c.billingPeriod
}

func (c *LedgerContract) BillingPrice() int64 {
	return c.billingPrice
}

func (c *LedgerContract) Drive() ID {
	return c.drive
}

func (c *LedgerContract) Owner() peer.ID {
	return c.owner
}

func (c *LedgerContract) ContractID() crypto.PubKey {
	return c.contractId
}

func (c *LedgerContract) PrivateKey() crypto.PrivKey {
	return c.privateKey
}

func (c *LedgerContract) Duration() int64 {
	return c.duration
}

func (c *LedgerContract) Root() cid.Cid {
	return c.root
}

func (c *LedgerContract) TotalSpace() int64 {
	return c.space
}

func (c *LedgerContract) Replicas() uint16 {
	return c.replicas
}

func (c *LedgerContract) MinReplicators() uint16 {
	return c.minReplicators
}

func (c *LedgerContract) PercentApprovers() uint8 {
	return c.percentApprovers
}

func (c *LedgerContract) MarshalBinary() ([]byte, error) {
	return MarshalLedgerContract(c)
}

func (c *LedgerContract) UnmarshalBinary(data []byte) error {
	basic, err := UnmarshalLedgerContract(data)
	if err != nil {
		return err
	}

	*c = *basic
	return nil
}

func (c *LedgerContract) MarshalJSON() ([]byte, error) {
	contractId, err := c.contractId.Bytes()
	if err != nil {
		return nil, err
	}

	return json.Marshal(&ledgerContractJSON{
		Drive:            c.drive,
		Owner:            c.owner,
		Duration:         c.duration,
		Root:             c.root,
		Space:            c.space,
		ContractId:       contractId,
		Replicas:         c.replicas,
		MinReplicators:   c.minReplicators,
		PercentApprovers: c.percentApprovers,
	})
}

func (c *LedgerContract) UnmarshalJSON(data []byte) error {
	cjson := &ledgerContractJSON{}
	err := json.Unmarshal(data, cjson)
	if err != nil {
		return err
	}

	contractId, err := crypto.UnmarshalPublicKey(cjson.ContractId)
	if err != nil {
		return err
	}
	c.drive = cjson.Drive
	c.owner = cjson.Owner
	c.duration = cjson.Duration
	c.root = cjson.Root
	c.space = cjson.Space
	c.contractId = contractId
	c.replicas = cjson.Replicas
	c.minReplicators = cjson.MinReplicators
	c.percentApprovers = cjson.PercentApprovers

	return nil
}
