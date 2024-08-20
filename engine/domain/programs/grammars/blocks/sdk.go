package blocks

import (
	"github.com/steve-care-software/webx/engine/domain/programs/grammars/blocks/lines"
	"github.com/steve-care-software/webx/engine/domain/programs/grammars/blocks/suites"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	return createBuilder()
}

// NewBlockBuilder creates a new block builder
func NewBlockBuilder() BlockBuilder {
	return createBlockBuilder()
}

// Builder represents a block list
type Builder interface {
	Create() Builder
	WithList(list []Block) Builder
	Now() (Blocks, error)
}

// Blocks represents blocks
type Blocks interface {
	List() []Block
	Fetch(name string) (Block, error)
}

// BlockBuilder represents a block builder
type BlockBuilder interface {
	Create() BlockBuilder
	WithName(name string) BlockBuilder
	WithLine(line lines.Line) BlockBuilder
	WithLines(lines lines.Lines) BlockBuilder
	WithSuites(suites suites.Suites) BlockBuilder
	Now() (Block, error)
}

// Block repreents a block
type Block interface {
	Name() string
	HasLine() bool
	Line() lines.Line
	HasLines() bool
	Lines() lines.Lines
	HasSuites() bool
	Suites() suites.Suites
}
