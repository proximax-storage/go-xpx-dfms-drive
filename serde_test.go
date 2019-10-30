package drive

import (
	"reflect"
	"testing"
)

func TestMarshalUnmarshalBinaryContract(t *testing.T) {
	in := RandContract(t)

	bytes, err := in.MarshalBinary()
	if err != nil {
		t.Fatal(err)
	}

	out := &Contract{}
	err = out.UnmarshalBinary(bytes)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(in, out) {
		t.Fatal("not equal")
	}
}

func TestMarshalUnmarshalJSONContract(t *testing.T) {
	in := RandContract(t)

	bytes, err := in.MarshalJSON()
	if err != nil {
		t.Fatal(err)
	}

	out := &Contract{}
	err = out.UnmarshalJSON(bytes)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(in, out) {
		t.Fatal("not equal")
	}
}

func TestMarshalUnmarshalInvite(t *testing.T) {
	in := RandInvite(t)

	bytes, err := MarshalInvite(in)
	if err != nil {
		t.Fatal(err)
	}

	out, err := UnmarshalInvite(bytes)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(in, out) {
		t.Fatal("not equal")
	}
}
