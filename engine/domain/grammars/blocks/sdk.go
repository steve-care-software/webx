package blocks

import (
	"github.com/steve-care-software/webx/engine/domain/grammars/blocks/lines"
)

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
	WithLines(lines lines.Lines) BlockBuilder
	Now() (Block, error)
}

// Block repreents a block
type Block interface {
	Name() string
	Lines() lines.Lines
}
