package selections

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// NewSelectionBuilder creates a new selection builder
func NewSelectionBuilder() SelectionBuilder {
	return createSelectionBuilder()
}

// NewChildBuilder creates a new child builder
func NewChildBuilder() ChildBuilder {
	return createChildBuilder()
}

// Builder represents a selections builder
type Builder interface {
	Create() Builder
	WithTreeName(treeName string) Builder
	WithList(list []Selection) Builder
	Now() (Selections, error)
}

// Selections represents selections
type Selections interface {
	TreeName() string
	List() []Selection
	Bytes() []byte
}

// SelectionBuilder represents a selection builder
type SelectionBuilder interface {
	Create() SelectionBuilder
	WithElementName(elementName string) SelectionBuilder
	WithList(list []Child) SelectionBuilder
	Now() (Selection, error)
}

// Selection represents a selection
type Selection interface {
	ElementName() string
	Bytes() []byte
	List() []Child
}

// ChildBuilder represents a child builder
type ChildBuilder interface {
	Create() ChildBuilder
	WithSelections(selections Selections) ChildBuilder
	WithContent(content []byte) ChildBuilder
	Now() (Child, error)
}

// Child represents a child
type Child interface {
	Bytes() []byte
	IsSelections() bool
	Selections() Selections
	IsContent() bool
	Content() []byte
}
