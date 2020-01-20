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

type SuperContract struct {
	// Identifier of SuperContract
	ID ID

	// Drive Contract
	Drive *drive.Contract

	// Supercontract file
	File cid.Cid

	// VM Functions list of Supercontract for Wasm VM 
	VMFunctionsList []*Function

	// Functions list of Supercontract
	FunctionsList []string
}

type Function struct {
	Name       string
	Parameters []int64
}
