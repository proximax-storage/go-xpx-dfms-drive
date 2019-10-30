package drive

import (
	"context"
)

const contractKey = "drive-contract"

func ContextWithContract(ctx context.Context, contract *Contract) context.Context {
	return context.WithValue(ctx, contractKey, contract)
}

func ContractFromContext(ctx context.Context) *Contract {
	value := ctx.Value(contractKey)
	if value == nil {
		return nil
	}

	d, ok := value.(*Contract)
	if !ok {
		panic("context value for contractKey is not a Drive")
	}

	return d
}
