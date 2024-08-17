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
	WithTokens(tokens []tokens.Token) Builder
	WithAsts(asts asts.AST) Builder
	Now() (Transpile, error)
}

// Transpile represents a transpile instance
type Transpile interface {
	Replacements() replacements.Replacements
	Tokens() []tokens.Token
	Asts() []asts.AST
	FetchToken(name string) (tokens.Token, error)
	FetchAST(name string) (asts.AST, error)
}
