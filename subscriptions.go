package drive

import (
	"context"
)

// InviteSubscription for DriveInvitations
type InviteSubscription interface {
	// Next waits and blocks till new Invitation is received
	Next(context.Context) (Invite, error)

	// Cancel stops subscription, like context canceling
	Cancel()
}

// ContractSubscription for Drive
type ContractSubscription interface {
	// Next waits and blocks till new Drive update received
	Next(context.Context) (Contract, error)

	// Cancel stops subscription, like context canceling
	Cancel()
}
