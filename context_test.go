package drive

import (
	"context"
	"reflect"
	"testing"
)

func TestContextWithContract(t *testing.T) {
	in := RandBasicContract(t)

	out := ContractFromContext(ContextWithContract(context.Background(), in))

	if !reflect.DeepEqual(in, out) {
		t.Fatal("not equal")
	}
}
