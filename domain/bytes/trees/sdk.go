package trees

import (
	"github.com/steve-care-software/syntax/domain/bytes/grammars"
	"github.com/steve-care-software/syntax/domain/bytes/grammars/values"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// NewBlockBuilder creates a new block builder
func NewBlockBuilder() BlockBuilder {
	return createBlockBuilder()
}

// NewLineBuilder creates a new line builder
func NewLineBuilder() LineBuilder {
	return createLineBuilder()
}

// NewElementsBuilder creates a new elements builder
func NewElementsBuilder() ElementsBuilder {
	return createElementsBuilder()
}

// NewElementBuilder creates a new element builder
func NewElementBuilder() ElementBuilder {
	return createElementBuilder()
}

// NewContentBuilder creates a new content builder
func NewContentBuilder() ContentBuilder {
	return createContentBuilder()
}

// Builder represents a tree builder
type Builder interface {
	Create() Builder
	WithGrammar(grammar grammars.Token) Builder
	WithBlock(block Block) Builder
	Now() (Tree, error)
}

// Tree represents a tree
type Tree interface {
	Grammar() grammars.Token
	Block() Block
}

// BlockBuilder represents a block builder
type BlockBuilder interface {
	Create() BlockBuilder
	WithLines(lines []Line) BlockBuilder
	Now() (Block, error)
}

// Block represents a block
type Block interface {
	Lines() []Line
	HasSuccessful() bool
	Successful() Line
}

// LineBuilder represents a line builder
type LineBuilder interface {
	Create() LineBuilder
	WithGrammar(grammar grammars.Line) LineBuilder
	WithElements(elements Elements) LineBuilder
	Now() (Line, error)
}

// Line represents a line of elements
type Line interface {
	Grammar() grammars.Line
	Elements() Elements
	IsSuccessful() bool
}

// ElementsBuilder represents elements builder
type ElementsBuilder interface {
	Create() ElementsBuilder
	WithList(list []Element) ElementsBuilder
	Now() (Elements, error)
}

// Elements represents elements
type Elements interface {
	List(isChannelsAccepted bool) []Element
}

// ElementBuilder represents an element builder
type ElementBuilder interface {
	Create() ElementBuilder
	WithGrammar(grammar grammars.Element) ElementBuilder
	WithContent(content Content) ElementBuilder
	WithAmount(amount uint) ElementBuilder
	IsChannel() ElementBuilder
	Now() (Element, error)
}

// Element represents an element
type Element interface {
	Grammar() grammars.Element
	Content() Content
	Amount() uint
	IsChannel() bool
}

// ContentBuilder represents a content builder
type ContentBuilder interface {
	Create() ContentBuilder
	WithValue(value values.Value) ContentBuilder
	WithTree(tree Tree) ContentBuilder
	Now() (Content, error)
}

// Content represents an element token
type Content interface {
	IsValue() bool
	Value() values.Value
	IsTree() bool
	Tree() Tree
}
