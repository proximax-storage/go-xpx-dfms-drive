package supercontract

import (
	"encoding/json"
	drive "github.com/proximax-storage/go-xpx-dfms-drive"

	"github.com/ipfs/go-cid"
	pb "github.com/proximax-storage/go-xpx-dfms-drive/pb"
)

func MarshalSuperContractProto(in *SuperContract) ([]byte, error) {
	driveContract, err := drive.MarshalContractProto(in.Drive)
	if err != nil {
		return nil, err
	}

	out := &pb.SContract{
		Id:        in.ID.Bytes(),
		Drive:     driveContract,
		File:      in.File.Bytes(),
		Vmversion: in.VMVersion,
		Functions: in.Functions,
	}

	return out.Marshal()
}

func UnmarshalSuperContractProto(data []byte) (*SuperContract, error) {
	in := new(pb.SContract)
	err := in.Unmarshal(data)
	if err != nil {
		return nil, err
	}

	out := &SuperContract{
		VMVersion: in.Vmversion,
		Functions: in.Functions,
	}

	out.ID, err = IDFromBytes(in.Id)
	if err != nil {
		return nil, err
	}

	out.Drive, err = drive.UnmarshalContractProto(in.Drive)
	if err != nil {
		return nil, err
	}

	out.File, err = cid.Cast(in.File)
	if err != nil {
		return nil, err
	}

	return out, nil
}

func MarshalSuperContractJSON(in *SuperContract) ([]byte, error) {
	out := &scContractJSON{
		Id:        in.ID.String(),
		Drive:     in.Drive,
		File:      in.File.String(),
		VMVersion: in.VMVersion,
		Functions: in.Functions,
	}

	return json.Marshal(out)
}

func UnmarshalSuperContractJSON(data []byte) (*SuperContract, error) {
	in := &scContractJSON{}
	err := json.Unmarshal(data, in)
	if err != nil {
		return nil, err
	}

	out := &SuperContract{
		Drive:     in.Drive,
		VMVersion: in.VMVersion,
		Functions: in.Functions,
	}

	out.ID, err = IDFromString(in.Id)
	if err != nil {
		return nil, err
	}

	out.File, err = cid.Decode(in.File)
	if err != nil {
		return nil, err
	}

	return out, nil
}

type scContractJSON struct {
	Id        string          `json:"id"`
	Drive     *drive.Contract `json:"drive"`
	File      string          `json:"file"`
	VMVersion uint64          `json:"vmversion"`
	Functions []string        `json:"functions"`
}
