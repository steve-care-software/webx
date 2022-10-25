package selections

import "github.com/steve-care-software/webx/domain/trees"

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// NewSelectionBuilder creates a new selection builder
func NewSelectionBuilder() SelectionBuilder {
	return createSelectionBuilder()
}

// NewElementBuilder creates a new element builder
func NewElementBuilder() ElementBuilder {
	return createElementBuilder()
}

// NewChildrenBuilder creates a new children builder
func NewChildrenBuilder() ChildrenBuilder {
	return createChildrenBuilder()
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
	WithElement(element Element) SelectionBuilder
	WithChildren(children Children) SelectionBuilder
	Now() (Selection, error)
}

// Selection represents a selection
type Selection interface {
	ElementName() string
	Bytes() []byte
	Content() Content
}

// Content represents a selection content
type Content interface {
	HasElement() bool
	Element() Element
	HasChildren() bool
	Children() Children
}

// ElementBuilder represents an element builder
type ElementBuilder interface {
	Create() ElementBuilder
	WithValue(value trees.Element) ElementBuilder
	IncludeChannelBytes() ElementBuilder
	Now() (Element, error)
}

// Element represents an element
type Element interface {
	Value() trees.Element
	IncludeChannelBytes() bool
	Bytes() []byte
}

// ChildrenBuilder represents a children builder
type ChildrenBuilder interface {
	Create() ChildrenBuilder
	WithElementName(elementName string) ChildrenBuilder
	WithList(list []Child) ChildrenBuilder
	Now() (Children, error)
}

// Children represents a children
type Children interface {
	ElementName() string
	Bytes() []byte
	List() []Child
}

// ChildBuilder represents a child builder
type ChildBuilder interface {
	Create() ChildBuilder
	WithSelections(selections Selections) ChildBuilder
	WithBytes(bytes []byte) ChildBuilder
	Now() (Child, error)
}

// Child represents a child
type Child interface {
	IsSelections() bool
	Selections() Selections
	IsBytes() bool
	Bytes() []byte
}
