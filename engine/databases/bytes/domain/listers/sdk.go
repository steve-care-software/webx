package listers

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents the lister builder
type Builder interface {
	Create() Builder
	WithKeyname(keyname string) Builder
	WithIndex(index uint64) Builder
	WithLength(length uint64) Builder
	Now() (Lister, error)
}

// Lister represents a lister
type Lister interface {
	Keyname() string
	Index() uint64
	Length() uint64
}
