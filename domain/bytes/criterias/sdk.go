package criterias

// Builder represents a criteria builder
type Builder interface {
	Create() Builder
	WithName(name string) Builder
	WithIndex(index uint) Builder
	WithRequirement(requirement uint) Builder
	WithChild(child Criteria) Builder
	Now() (Criteria, error)
}

// Criteria represents a criteria
type Criteria interface {
	Name() string
	Index() bool
	HasRequirement() bool
	Requirement() *uint
	HasChild() bool
	Child() Criteria
}
