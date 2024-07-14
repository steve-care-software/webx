package amounts

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents the amount instruction
type Builder interface {
	Create() Builder
	WithContext(context string) Builder
	WithReturn(ret string) Builder
	Now() (Amount, error)
}

// Amount represents an amount
type Amount interface {
	Context() string
	Return() string
}
