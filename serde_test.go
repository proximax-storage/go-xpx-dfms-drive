package drive

import (
	"reflect"
	"testing"
)

func TestMarshalUnmarshalBasicContract(t *testing.T) {
	in := RandBasicContract(t)

	bytes, err := MarshalBasicContract(in)
	if err != nil {
		t.Fatal(err)
	}

	out, err := UnmarshalBasicContract(bytes)
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
