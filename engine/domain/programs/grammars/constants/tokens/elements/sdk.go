package elements

// NewBuilder creates a new builder
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents the element builder
type Builder interface {
	Create() Builder
	WithRule(rule string) Builder
	WithConstant(constant string) Builder
	Now() (Element, error)
}

// Element represents a constant element
type Element interface {
	IsRule() bool
	Rule() string
	IsConstant() bool
	Constant() string
}
