package grammars

import (
	"github.com/steve-care-software/webx/engine/domain/programs/grammars/blocks"
	"github.com/steve-care-software/webx/engine/domain/programs/grammars/blocks/lines/tokens/elements"
	"github.com/steve-care-software/webx/engine/domain/programs/grammars/rules"
	"github.com/steve-care-software/webx/engine/domain/programs/grammars/syscalls"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents the grammar builder
type Builder interface {
	Create() Builder
	WithVersion(version uint) Builder
	WithRoot(root elements.Element) Builder
	WithRules(rules rules.Rules) Builder
	WithBlocks(blocks blocks.Blocks) Builder
	WithSyscalls(syscalls syscalls.Syscalls) Builder
	WithOmissions(omissions elements.Elements) Builder
	Now() (Grammar, error)
}

// Grammar represents a grammar
type Grammar interface {
	Version() uint
	Root() elements.Element
	Rules() rules.Rules
	Blocks() blocks.Blocks
	HasSyscalls() bool
	Syscalls() syscalls.Syscalls
	HasOmissions() bool
	Omissions() elements.Elements
}
