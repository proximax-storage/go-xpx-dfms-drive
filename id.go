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

// NewIDFromInvite returns new ID generated from Invite
func NewIDFromInvite(invite Invite) (ID, error) {
	data, err := MarshalInvite(invite)
	if err != nil {
		return nil, err
	}

	return NewIDFromBytes(data)
}

// NewIDFromBytes creates new ID by hashing given Data
func NewIDFromBytes(data []byte) (ID, error) {
	return crypto.UnmarshalEd25519PublicKey(data)
}

// IdFromString creates new ID by hashing given Data
func NewIDFromString(data string) (ID, error) {
	return crypto.UnmarshalEd25519PublicKey([]byte(data))
}

// IdFromBytes creates new ID by hashing given Data
func IdFromBytes(data []byte) (ID, error) {
	return crypto.UnmarshalPublicKey(data)
}

// NewIDFromBytes creates new ID by hashing given Data
func IdToBytes(id ID) ([]byte, error) {
	data, err := crypto.MarshalPublicKey(id)
	if err != nil {
		return nil, err
	}
	return data, nil
}
