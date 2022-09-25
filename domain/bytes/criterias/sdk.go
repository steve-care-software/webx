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
	WithRequirement(requirement []byte) Builder
	WithChild(child Criteria) Builder
	Now() (Criteria, error)
}

// Criteria represents a criteria
type Criteria interface {
	Name() string
	Index() uint
	Content() Content
}

// Content represents a criteria's content
type Content interface {
	IsRequirement() bool
	Requirement() []byte
	IsChild() bool
	Child() Criteria
}
