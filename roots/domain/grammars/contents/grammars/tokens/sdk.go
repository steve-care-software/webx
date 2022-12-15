package tokens

import (
	"github.com/steve-care-software/webx/roots/domain/blockchains/cryptography/hash"
)

const minTokenSize = hash.Size + 8

// NewAdapter creates a new adapter for tests
func NewAdapter() Adapter {
	hashAdapter := hash.NewAdapter()
	linesAdapter := NewLinesAdapter()
	builder := NewBuilder()
	return createAdapter(hashAdapter, linesAdapter, builder)
}

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// NewLinesAdapter creates a new lines adapter
func NewLinesAdapter() LinesAdapter {
	builder := NewLinesBuilder()
	lineAdapter := NewLineAdapter()
	return createLinesAdapter(builder, lineAdapter)
}

// NewLinesBuilder creates a new lines Builder
func NewLinesBuilder() LinesBuilder {
	return createLinesBuilder()
}

// NewLineAdapter creates a new line adapter
func NewLineAdapter() LineAdapter {
	hashAdapter := hash.NewAdapter()
	builder := NewLineBuilder()
	return createLineAdapter(hashAdapter, builder)
}

// NewLineBuilder creates a new line builder
func NewLineBuilder() LineBuilder {
	return createLineBuilder()
}

// Adapter represents a token adapter
type Adapter interface {
	ToContent(ins Token) ([]byte, error)
	ToToken(content []byte) (Token, error)
}

// Builder represents a token builder
type Builder interface {
	Create() Builder
	WithHash(hash hash.Hash) Builder
	WithLines(lines Lines) Builder
	Now() (Token, error)
}

// Token represents token metadata
type Token interface {
	Hash() hash.Hash
	Lines() Lines
}

// LinesAdapter represents a lines adapter
type LinesAdapter interface {
	ToContent(ins Lines) ([]byte, error)
	ToLines(content []byte) (Lines, error)
}

// LinesBuilder represents lines builder
type LinesBuilder interface {
	Create() LinesBuilder
	WithList(list []Line) LinesBuilder
	Now() (Lines, error)
}

// Lines represents lines
type Lines interface {
	List() []Line
}

// LineAdapter represents a line adapter
type LineAdapter interface {
	ToContent(ins Line) ([]byte, error)
	ToLine(content []byte) (Line, error)
}

// LineBuilder represents a line builder
type LineBuilder interface {
	Create() LineBuilder
	WithElements(elements []hash.Hash) LineBuilder
	Now() (Line, error)
}

// Line represents a token line
type Line interface {
	Elements() []hash.Hash
}
