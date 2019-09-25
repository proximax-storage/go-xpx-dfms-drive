package drive

import (
	"encoding"

	"github.com/ipfs/go-cid"
	"github.com/libp2p/go-libp2p-core/peer"
)

// Contract is an agreement between client and replicator peers on some amount of disk space
type Contract interface {
	encoding.BinaryMarshaler
	encoding.BinaryUnmarshaler

	// Drive returns identifier of Contract's Drive
	Drive() ID

	// Owner is a peer which owns the Drive
	Owner() peer.ID

	// Members are peers responsible for Drive
	Members() []peer.ID

	// Duration of the contract.
	// NOTE: May return time in milliseconds or amount of blocks,
	// depending on the implementation.
	Duration() uint64

	// Created is a Drive starting moment.
	// Usually derived from ab invitation
	// NOTE: May return time in milliseconds or amount of blocks,
	// depending on the implementation.
	Created() uint64

	// Root is a drive of Drive's directory
	// Formerly, represents the identifier of all content stored on drive.
	Root() cid.Cid

	// TotalSpace specifies total physical space used by Drive
	// on member nodes.
	TotalSpace() uint64
}
