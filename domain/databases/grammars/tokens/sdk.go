package tokens

import (
	"github.com/steve-care-software/webx/domain/cryptography/hash"
)

const minSuiteSize = 2

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

// NewSuitesAdapter creates a new suites adapter
func NewSuitesAdapter() SuitesAdapter {
	builder := NewSuitesBuilder()
	suiteAdapter := NewSuiteAdapter()
	return createSuitesAdapter(builder, suiteAdapter)
}

// NewSuitesBuilder creates a new suites builder
func NewSuitesBuilder() SuitesBuilder {
	return createSuitesBuilder()
}

// NewSuiteAdapter creates a new suite adapter
func NewSuiteAdapter() SuiteAdapter {
	builder := NewSuiteBuilder()
	return createSuiteAdapter(builder)
}

// NewSuiteBuilder creates a new suite builder instance
func NewSuiteBuilder() SuiteBuilder {
	return createSuiteBuilder()
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
	WithSuites(suites Suites) Builder
	Now() (Token, error)
}

// Token represents token metadata
type Token interface {
	Hash() hash.Hash
	Lines() Lines
	HasSuites() bool
	Suites() Suites
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

// SuitesAdapter represents a suites adapter
type SuitesAdapter interface {
	ToContent(ins Suites) ([]byte, error)
	ToSuites(content []byte) (Suites, error)
}

// SuitesBuilder represents a suites builder
type SuitesBuilder interface {
	Create() SuitesBuilder
	WithList(list []Suite) SuitesBuilder
	Now() (Suites, error)
}

// Suites represents suites
type Suites interface {
	List() []Suite
}

// SuiteAdapter represents a suite adapter
type SuiteAdapter interface {
	ToContent(ins Suite) ([]byte, error)
	ToSuite(content []byte) (Suite, error)
}

// SuiteBuilder represents a suite builder
type SuiteBuilder interface {
	Create() SuiteBuilder
	WithContent(content []byte) SuiteBuilder
	IsValid() SuiteBuilder
	Now() (Suite, error)
}

// Suite represents a suite
type Suite interface {
	IsValid() bool
	Content() []byte
}
