package lexers

// NewBuilder creates a new builder
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents the lexer builder
type Builder interface {
	Create() Builder
	WithOutput(output []byte) Builder
	IsFail() Builder
	Now() (Lexer, error)
}

// Lexer represents a suite lexer
type Lexer interface {
	Output() []byte
	IsFail() bool
}
