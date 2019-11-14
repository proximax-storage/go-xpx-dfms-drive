package drive

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContextWithContract(t *testing.T) {
	in := RandContract(t)
	out := ContractFromContext(ContextWithContract(context.Background(), in))
	assert.Equal(t, in, out)
}
