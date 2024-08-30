package constants

import "github.com/steve-care-software/webx/engine/domain/programs/grammars/constants/tokens"

// Builder represents the constants builder
type Builder interface {
	Create() Builder
	WithList(list []Constant) Builder
	Now() (Constants, error)
}

// Constants represents constants
type Constants interface {
	List() []Constant
}

// ConstantBuilder represents the constant builder
type ConstantBuilder interface {
	Create() ConstantBuilder
	WithName(name string) ConstantBuilder
	WithTokens(tokens tokens.Tokens) ConstantBuilder
	Now() (Constant, error)
}

// Constant represents a constant
type Constant interface {
	Name() string
	Tokens() tokens.Tokens
}
