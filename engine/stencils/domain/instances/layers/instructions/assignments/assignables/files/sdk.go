package files

import (
	"github.com/steve-care-software/webx/engine/states/domain/hash"
	"github.com/steve-care-software/webx/engine/stencils/domain/instances/layers/instructions/assignments/assignables/files/opens"
	"github.com/steve-care-software/webx/engine/stencils/domain/instances/layers/instructions/assignments/assignables/files/reads"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// Builder represents a file builder
type Builder interface {
	Create() Builder
	WithOpen(open opens.Open) Builder
	WithRead(read reads.Read) Builder
	WithLength(length string) Builder
	WithExists(exists string) Builder
	Now() (File, error)
}

// File represents a file
type File interface {
	Hash() hash.Hash
	IsOpen() bool
	Open() opens.Open
	IsRead() bool
	Read() reads.Read
	IsExists() bool
	Exists() string
	IsLength() bool
	Length() string
}
