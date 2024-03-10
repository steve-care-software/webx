package reverts

// Revert represents a revert
type Revert interface {
	HasIndex() bool
	Index() uint
}
