package supercontract

import (
	"github.com/ipfs/go-cid"
	"github.com/multiformats/go-multihash"
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
	Hash multihash.Multihash
	Name string
}

func (f *Function) SetHash(prefix string) error {
	var err error
	f.Hash, err = multihash.Sum([]byte(prefix+f.Name), multihash.SHA2_256, 32)
	return err
}
