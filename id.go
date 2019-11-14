package drive

import (
	"github.com/ipfs/go-cid"
	"github.com/libp2p/go-libp2p-core/crypto"
)

const (
	Codec     = 0xC
	CodecName = "drive"
)

func init() {
	cid.Codecs[CodecName] = Codec
	cid.CodecToStr[Codec] = CodecName
}

// ID represents identifier of the drive
type ID = crypto.PubKey
