package drive

// Billing represent Contract Billing info for Contract billing history
// TODO: [issue: #13] that it will be expanded with other useful values
type Billing struct {
	StartBlock int64
	EndBlock   int64
}

// IsCurrentPeriod checks if given block was in billing period
func (b *Billing) IsCurrentPeriod(height int64) bool {
	return (b.StartBlock <= height) && (height <= b.EndBlock)
}
