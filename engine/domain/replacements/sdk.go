package replacements

import (
	"github.com/steve-care-software/webx/engine/domain/asts"
	"github.com/steve-care-software/webx/engine/domain/grammars/tokens"
)

// Builder represents a builder
type Builder interface {
	Create() Builder
	WithList(list []Replacement) Builder
	Now() (Replacements, error)
}

// Replacements represents a replacement list
type Replacements interface {
	List() []Replacement
}

// ReplacementBuilder represents a replacement builder
type ReplacementBuilder interface {
	Create() ReplacementBuilder
	WithOrigin(origin tokens.Token) ReplacementBuilder
	WithTarget(target asts.AST) ReplacementBuilder
	Now() (Replacement, error)
}

// Replacement represents a replacement
type Replacement interface {
	Origin() tokens.Token
	Target() asts.AST
}
