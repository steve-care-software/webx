package uncovers

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// NewUncoverBuilder creates a new uncover builder
func NewUncoverBuilder() UncoverBuilder {
	return createUncoverBuilder()
}

// NewLineBuilder creates a new line builder
func NewLineBuilder() LineBuilder {
	return createLineBuilder()
}

// Builder represents uncovers builder
type Builder interface {
	Create() Builder
	WithList(list []Uncover) Builder
	Now() (Uncovers, error)
}

// Uncovers represents uncovers
type Uncovers interface {
	List() []Uncover
}

// UncoverBuilder represents an uncover builder
type UncoverBuilder interface {
	Create() UncoverBuilder
	WithName(name string) UncoverBuilder
	WithLine(line Line) UncoverBuilder
	Now() (Uncover, error)
}

// Uncover represents an uncover element
type Uncover interface {
	Name() string
	Line() Line
}

// LineBuilder represents a line builder
type LineBuilder interface {
	Create() LineBuilder
	WithIndex(index uint) LineBuilder
	WithElements(elements []string) LineBuilder
	Now() (Line, error)
}

// Line represents a line
type Line interface {
	Index() uint
	Elements() []string
}
