package supercontract

import (
	drive "github.com/proximax-storage/go-xpx-dfms-drive"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMarshalUnmarshalBinarySuperContract(t *testing.T) {
	in := randSuperContract(t)

	data, err := in.MarshalBinary()
	require.Nil(t, err, err)

	out := &SuperContract{}
	err = out.UnmarshalBinary(data)
	require.Nil(t, err, err)

	assert.Equal(t, in, out)
}

func TestMarshalUnmarshalJSONSuperContract(t *testing.T) {
	in := randSuperContract(t)

	data, err := in.MarshalJSON()
	require.Nil(t, err, err)

	out := &SuperContract{}
	err = out.UnmarshalJSON(data)
	require.Nil(t, err, err)

	assert.Equal(t, in, out)
}

func randSuperContract(t *testing.T) *SuperContract {
	return &SuperContract{
		ID:        RandID(t),
		Drive:     drive.RandContract(t),
		File:      randCID(t),
		VMVersion: 3,
		Functions: []string{"main"},
		Status:    DeactivatedByDriveEnd,
	}
}
