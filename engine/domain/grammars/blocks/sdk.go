package blocks

import (
	"github.com/steve-care-software/webx/engine/domain/grammars/blocks/lines"
)

// Builder represents a block builder
type Builder interface {
	Create() Builder
	WithName(name string) Builder
	WithLines(lines lines.Lines) Builder
	Now() (Block, error)
}

// Block repreents a block
type Block interface {
	Name() string
	Lines() lines.Lines
}
