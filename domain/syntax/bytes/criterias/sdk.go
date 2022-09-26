package criterias

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents a criteria builder
type Builder interface {
	Create() Builder
	WithName(name string) Builder
	WithIndex(index uint) Builder
	WithChild(child Criteria) Builder
	IncludeChannels() Builder
	Now() (Criteria, error)
}

// Criteria represents a criteria
type Criteria interface {
	Name() string
	Index() uint
	IncludeChannels() bool
	HasChild() bool
	Child() Criteria
}
