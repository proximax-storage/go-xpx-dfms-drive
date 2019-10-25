package drive

import (
	"encoding"
	"encoding/json"

	"github.com/ipfs/go-cid"
	"github.com/libp2p/go-libp2p-core/crypto"
	"github.com/libp2p/go-libp2p-core/peer"
)

// TODO Currently Contract has two ids two support Blockchain and non-Blockchain DFMS versions. Merge them into one
// Contract is an agreement between client and replicator peers on some amount of disk space
type Contract interface {
	encoding.BinaryMarshaler
	encoding.BinaryUnmarshaler
	json.Marshaler
	json.Unmarshaler

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

	// PrivateKey specifies private key for signing
	PrivateKey() crypto.PrivKey

	ContractID() crypto.PubKey

	Replicas() int8

	MinReplicatorsDelta() int8

	MinApproversDelta() int8
}
