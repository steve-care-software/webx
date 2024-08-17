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
	WithAsts(asts []asts.AST) Builder
	WithBlocks(blocks []blocks.Block) Builder
	WithTokens(tokens []tokens.Token) Builder
	WithRules(rules []rules.Rule) Builder
	Now() (Grammar, error)
}

// Grammar represents a grammar
type Grammar interface {
	BlockEntry() string
	Asts() []asts.AST
	Blocks() []blocks.Block
	Tokens() []tokens.Token
	Rules() []rules.Rule
	FetchBlock(name string) (blocks.Block, error)
	FetchToken(name string) (tokens.Token, error)
	FetchRule(name string) (rules.Rule, error)
	HasOmissions() bool
	Omissions() []string
}
