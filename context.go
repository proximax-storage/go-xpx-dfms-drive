package drive

import (
	"context"
)

const contractKey = "drive-contract"

func ContextWithContract(ctx context.Context, contract *BasicContract) context.Context {
	return context.WithValue(ctx, contractKey, contract)
}

func ContractFromContext(ctx context.Context) *BasicContract {
	value := ctx.Value(contractKey)
	if value == nil {
		return nil
	}

	d, ok := value.(BasicContract)
	if !ok {
		panic("context value for contractKey is not a Drive")
	}

	return &d
}
