package transpiles

import (
	"github.com/steve-care-software/webx/engine/domain/asts"
	"github.com/steve-care-software/webx/engine/domain/grammars/blocks/lines/replacements"
	"github.com/steve-care-software/webx/engine/vms/domain/instances/routes/tokens"
)

// Builder represents a transpile builder
type Builder interface {
	Create() Builder
	WithReplacements(replacements replacements.Replacements) Builder
	WithTokens(tokens tokens.Tokens) Builder
	WithAsts(asts asts.ASTs) Builder
	Now() (Transpile, error)
}

// Transpile represents a transpile instance
type Transpile interface {
	Replacements() replacements.Replacements
	Tokens() tokens.Tokens
	Asts() asts.ASTs
}