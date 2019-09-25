package drive

import (
	"github.com/ipfs/go-cid"
	"github.com/multiformats/go-multihash"
)

const (
	Codec     = 0xC
	CodecName = "drive"
)

var NilID = ID{}

func init() {
	cid.Codecs[CodecName] = Codec
	cid.CodecToStr[Codec] = CodecName
}

// ID represents identifier of the drive
type ID cid.Cid

// NewIDFromBytes creates new ID by hashing given Data
func NewIDFromBytes(data []byte) (ID, error) {
	hash, err := multihash.Sum(data, multihash.SHA2_256, -1)
	if err != nil {
		return ID{}, nil
	}

	return IDFromCID(cid.NewCidV1(Codec, hash)), nil
}

// IDFromCID casts cid to ID
func IDFromCID(id cid.Cid) ID {
	return ID(id)
}

// CID casts ID to CID
func (id ID) CID() cid.Cid {
	return cid.Cid(id)
}

func (id *ID) MarshalJSON() ([]byte, error) {
	return id.CID().MarshalJSON()
}

func (id *ID) UnmarshalJSON(data []byte) error {
	cid := cid.Cid{}
	err := cid.UnmarshalJSON(data)
	if err != nil {
		return err
	}

	*id = IDFromCID(cid)
	return nil
}

func (id *ID) MarshalBinary() ([]byte, error) {
	return MarshalID(*id)
}

func (id *ID) UnmarshalBinary(data []byte) (err error) {
	*id, err = UnmarshalID(data)
	return
}
