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
	WithMatch(match []byte) Builder
	IncludeChannels() Builder
	Now() (Criteria, error)
}

// Criteria represents a criteria
type Criteria interface {
	Name() string
	Index() uint
	IncludeChannels() bool
	HasContent() bool
	Content() Content
}

// Content represents a criteria content
type Content interface {
	IsChild() bool
	Child() Criteria
	IsMatch() bool
	Match() []byte
}
