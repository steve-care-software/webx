package elements

// NewBuilder creates a new builder
func NewBuilder() Builder {
	return createBuilder()
}

// NewElementBuilder creates a new element builder
func NewElementBuilder() ElementBuilder {
	return createElementBuilder()
}

// Builder represents an elments list
type Builder interface {
	Create() Builder
	WithList(list []Element) Builder
	Now() (Elements, error)
}

// Elements represents elements
type Elements interface {
	List() []Element
	Fetch(name string) (Element, error)
}

// ElementBuilder represents the element builder
type ElementBuilder interface {
	Create() ElementBuilder
	WithRule(rule string) ElementBuilder
	WithBlock(block string) ElementBuilder
	WithSpacer(spacer string) ElementBuilder
	WithConstant(constant string) ElementBuilder
	Now() (Element, error)
}

// Element represents an element
type Element interface {
	Name() string
	IsRule() bool
	Rule() string
	IsBlock() bool
	Block() string
	IsSpacer() bool
	Spacer() string
	IsConstant() bool
	Constant() string
}
