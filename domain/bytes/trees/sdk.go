package trees

import (
	"github.com/steve-care-software/syntax/domain/bytes/grammars"
	"github.com/steve-care-software/syntax/domain/bytes/grammars/values"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// NewTreeBuilder creates a new tree builder instance
func NewTreeBuilder() TreeBuilder {
	return createTreeBuilder()
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

// Builder represents a trees builder
type Builder interface {
	Create() Builder
	WithList(list []Tree) Builder
	Now() (Trees, error)
}

// Trees represents a trees
type Trees interface {
	List() []Tree
}

// TreeBuilder represents a tree builder
type TreeBuilder interface {
	Create() TreeBuilder
	WithGrammar(grammar grammars.Token) TreeBuilder
	WithBlock(block Block) TreeBuilder
	WithSuffix(suffix Trees) TreeBuilder
	Now() (Tree, error)
}

// Tree represents a tree
type Tree interface {
	Grammar() grammars.Token
	Block() Block
	HasSuffix() bool
	Suffix() Trees
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
	List() []Element
}

// ElementBuilder represents an element builder
type ElementBuilder interface {
	Create() ElementBuilder
	WithGrammar(grammar grammars.Element) ElementBuilder
	WithContent(content Content) ElementBuilder
	WithAmount(amount uint) ElementBuilder
	Now() (Element, error)
}

// Element represents an element
type Element interface {
	Grammar() grammars.Element
	Content() Content
	Amount() uint
}

// ContentBuilder represents a content builder
type ContentBuilder interface {
	Create() ContentBuilder
	WithValue(value Value) ContentBuilder
	WithTree(tree Tree) ContentBuilder
	Now() (Content, error)
}

// Content represents an element token
type Content interface {
	IsValue() bool
	Value() Value
	IsTree() bool
	Tree() Tree
}

// ValueBuilder represents a value builder
type ValueBuilder interface {
	Create() ValueBuilder
	WithContent(content values.Value) ValueBuilder
	WithPrefix(prefix Trees) ValueBuilder
	Now() (Value, error)
}

// Value represents a value
type Value interface {
	Content() values.Value
	HasPrefix() bool
	Prefix() Trees
}
