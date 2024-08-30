package elements

// Builder represents the element builder
type Builder interface {
	Create() Builder
	WithRule(rule string) Builder
	WithSpacer(spacer string) Builder
	Now() (Element, error)
}

// Element represents a spacer element
type Element interface {
	IsRule() bool
	Rule() string
	IsSpacer() bool
	Spacer() string
}
