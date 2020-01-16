package drive

import (
	"github.com/ipfs/go-cid"
	"github.com/libp2p/go-libp2p-core/crypto"
)

type ExecutionStatus uint8

const (
	ExecutionUnknown ExecutionStatus = iota
	ExecutionStarted
	ExecutionSuccess
	ExecutionFailure
)

type ScID crypto.PubKey

type SuperContract struct {
	// Identifier of Supercontract
	ID ScID

	// Drive Contract
	Drive *Contract

	// Supercontract file
	File cid.Cid

	// VM Functions list of Supercontract for Wasm VM 
	VMFunctionsList []string

	// Functions list of Supercontract
	FunctionsList []string
}
