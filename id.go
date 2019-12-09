package drive

import (
	"encoding/hex"
	"encoding/json"
	"errors"

	"github.com/ipfs/go-cid"
	"github.com/libp2p/go-libp2p-core/crypto"
	"github.com/multiformats/go-multihash"
)

// Use NilID instead of "nil"
var NilID = ID{}

var ErrWrongID = errors.New("wrong drive.ID")

// ID represents identifier of the drive
type ID cid.Cid

func (id ID) CID() cid.Cid {
	return (cid.Cid)(id)
}

func (id ID) String() string {
	return ((cid.Cid)(id)).String()
}

func (id ID) Equals(id2 ID) bool {
	return id == id2
}

func (id ID) Loggable() map[string]interface{} {
	return map[string]interface{}{
		"drive": id.String(),
	}
}

func (id ID) Bytes() []byte {
	return ((cid.Cid)(id)).Bytes()
}

func (id ID) MarshalBinary() ([]byte, error) {
	return id.Bytes(), nil
}

func (id *ID) UnmarshalBinary(data []byte) (err error) {
	*id, err = IDFromBytes(data)
	return
}

func (id ID) MarshalJSON() ([]byte, error) {
	return json.Marshal(((cid.Cid)(id)).String())
}

func (id *ID) UnmarshalJSON(data []byte) (err error) {
	var v string
	err = json.Unmarshal(data, &v)
	if err != nil {
		return
	}

	*id, err = IDFromString(v)
	return
}

// IDFromString decodes string into ID
// Also accepts hex encoded public key representation.
// TODO Use own crypto repository
func IDFromString(str string) (ID, error) {
	id, err := cid.Decode(str)
	if err == nil {
		if id.Type() != codec {
			return NilID, ErrWrongID
		}

		return ID(id), nil
	}

	data, err := hex.DecodeString(str)
	if err != nil {
		return NilID, ErrWrongID
	}

	key, err := crypto.UnmarshalPublicKey(data)
	if err != nil {
		return NilID, ErrWrongID
	}

	return IDFromPubKey(key)
}

func IDFromBytes(data []byte) (ID, error) {
	id, err := cid.Cast(data)
	if err != nil {
		return NilID, ErrWrongID
	}

	if id.Type() != codec {
		return NilID, ErrWrongID
	}

	return ID(id), nil
}

func IDFromPubKey(key crypto.PubKey) (ID, error) {
	b, err := key.Bytes()
	if err != nil {
		return NilID, err
	}

	mh, _ := multihash.Sum(b, multihash.IDENTITY, -1)
	return ID(cid.NewCidV1(codec, mh)), nil
}

func IDToPubKey(id ID) (crypto.PubKey, error) {
	dh, err := multihash.Decode(id.CID().Hash())
	if err != nil {
		return nil, err
	}

	if dh.Code != multihash.IDENTITY {
		return nil, ErrWrongID
	}

	return crypto.UnmarshalPublicKey(dh.Digest)
}

const (
	// TODO Register in multicodec table: https://github.com/multiformats/multicodec/blob/master/table.csv
	codec     = 0xC
	codecName = "proximax-drive"
)

func init() {
	cid.Codecs[codecName] = codec
	cid.CodecToStr[codec] = codecName
}
