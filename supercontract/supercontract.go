package supercontract

import (
	"github.com/ipfs/go-cid"
	drive "github.com/proximax-storage/go-xpx-dfms-drive"
)

type ExecutionStatus uint8

const (
	ExecutionUnknown ExecutionStatus = iota
	ExecutionStarted
	ExecutionSuccess
	ExecutionFailure
)

type Status uint8

const (
	Active                   Status = 0
	DeactivatedByParticipant        = 100
	DeactivatedByDriveEnd           = 101
)

type SuperContract struct {
	// Identifier of SuperContract
	ID ID

	// Drive Contract
	Drive *drive.Contract

	// Supercontract file
	File cid.Cid

	// Version of Wasm VM
	VMVersion uint64

	// SuperContract functions
	Functions []string

	// SuperContract status
	Status Status
}

type Function struct {
	Name       string
	Parameters []int64
}

func (c *SuperContract) MarshalBinary() ([]byte, error) {
	return MarshalSuperContractProto(c)
}

func (c *SuperContract) UnmarshalBinary(data []byte) error {
	out, err := UnmarshalSuperContractProto(data)
	if err != nil {
		return err
	}

	*c = *out
	return nil
}

func (c *SuperContract) MarshalJSON() ([]byte, error) {
	return MarshalSuperContractJSON(c)
}

func (c *SuperContract) UnmarshalJSON(data []byte) error {
	out, err := UnmarshalSuperContractJSON(data)
	if err != nil {
		return err
	}

	*c = *out
	return nil
}
