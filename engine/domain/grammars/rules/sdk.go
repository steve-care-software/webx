package rules

// Builder represents a rule builder
type Builder interface {
	Create() Builder
	WithName(name string) Builder
	WithBytes(bytes []byte) Builder
	Now() (Rule, error)
}

// Rule represents a rule
type Rule interface {
	Name() string
	Bytes() []byte
}
