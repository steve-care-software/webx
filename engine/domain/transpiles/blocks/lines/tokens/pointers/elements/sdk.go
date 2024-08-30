package elements

// Builder represents the element builder
type Builder interface {
	Create() Builder
	WithToken(token string) Builder
	WithRule(rule string) Builder
	Now() (Element, error)
}

// Element represents an element
type Element interface {
	IsToken() bool
	Token() string
	IsRule() bool
	Rule() string
}
