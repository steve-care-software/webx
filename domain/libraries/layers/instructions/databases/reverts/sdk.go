package reverts

// Builder represents a revert builder
type Builder interface {
	Create() Builder
	WithIndex(index string) Builder
	Now() (Revert, error)
}

// Revert represents a revert
type Revert interface {
	HasIndex() bool
	Index() string
}
