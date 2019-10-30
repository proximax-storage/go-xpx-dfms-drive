package drive

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
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

	assert.Equal(t, in.Drive, out.Drive)
	assert.Equal(t, in.Owner, out.Owner)
	assert.Equal(t, in.State, out.State)
	assert.Equal(t, in.Space, out.Space)
	assert.Equal(t, in.Members, out.Members)
	assert.Equal(t, in.ContractId, out.ContractId)
	assert.Equal(t, in.Duration, out.Duration)
	assert.Equal(t, in.Root, out.Root)
	assert.Equal(t, in.MinReplicators, out.MinReplicators)
	assert.Equal(t, in.PercentApprovers, out.PercentApprovers)
	assert.Equal(t, in.BillingPrice, out.BillingPrice)
	assert.Equal(t, in.BillingPeriod, out.BillingPeriod)

	assert.Nil(t, out.PrivateKey)
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

	assert.Equal(t, in.Drive, out.Drive)
	assert.Equal(t, in.Owner, out.Owner)
	assert.Equal(t, in.State, out.State)
	assert.Equal(t, in.Space, out.Space)
	assert.Equal(t, in.Members, out.Members)
	assert.Equal(t, in.ContractId, out.ContractId)
	assert.Equal(t, in.Duration, out.Duration)
	assert.Equal(t, in.Root, out.Root)
	assert.Equal(t, in.MinReplicators, out.MinReplicators)
	assert.Equal(t, in.PercentApprovers, out.PercentApprovers)
	assert.Equal(t, in.BillingPrice, out.BillingPrice)
	assert.Equal(t, in.BillingPeriod, out.BillingPeriod)
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
