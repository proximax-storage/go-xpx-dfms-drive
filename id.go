package drive

import (
	"github.com/ipfs/go-cid"
	"github.com/multiformats/go-multihash"
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
type ID = cid.Cid

// NewIDFromBytes creates new ID by hashing given Data
func NewIDFromBytes(data []byte) (ID, error) {
	hash, err := multihash.Sum(data, multihash.SHA2_256, -1)
	if err != nil {
		return ID{}, nil
	}

	return cid.NewCidV1(Codec, hash), nil
}
