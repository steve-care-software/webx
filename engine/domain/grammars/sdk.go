package grammars

import (
	"github.com/steve-care-software/webx/engine/domain/asts"
	"github.com/steve-care-software/webx/engine/domain/grammars/blocks"
	"github.com/steve-care-software/webx/engine/domain/grammars/rules"
	"github.com/steve-care-software/webx/engine/domain/grammars/tokens"
)

// Builder represents the grammar builder
type Builder interface {
	Create() Builder
	WithBlockEntry(blockEntry string) Builder
	WithOmissions(omissions []string) Builder
	WithAsts(asts asts.ASTs) Builder
	WithBlocks(blocks blocks.Blocks) Builder
	WithTokens(tokens tokens.Tokens) Builder
	Now() (Grammar, error)
}

// Grammar represents a grammar
type Grammar interface {
	BlockEntry() string
	Asts() asts.ASTs
	Blocks() blocks.Blocks
	Tokens() tokens.Tokens
	Rules() rules.Rules
	HasOmissions() bool
	Omissions() []string
}
