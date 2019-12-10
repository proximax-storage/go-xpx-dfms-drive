package drive

type Height int64

// Billing represent Contract Billing info for Contract billing history
type Billing struct {
	Start Height
	End   Height
}

// IsCurrentPeriod check is current billing period
func (b *Billing) IsCurrentPeriod(height Height) bool {
	return (b.Start <= height) && (height <= b.End)
}
