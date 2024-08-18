package elements

// NewBuilder creates a new builder
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents the element builder
type Builder interface {
	Create() Builder
	WithRule(rule string) Builder
	WithBlock(block string) Builder
	Now() (Element, error)
}

// Element represents an element
type Element interface {
	IsRule() bool
	Rule() string
	IsBlock() bool
	Block() string
}
