package trees

import "github.com/steve-care-software/webx/domain/databases/entities"

// Builder represents a trees builder
type Builder interface {
	Create() Trees
	WithEntity(entity entities.Entity) Trees
	WithList(list []Tree) Trees
	Now() (Trees, error)
}

// Trees represents trees
type Trees interface {
	Entity() entities.Entity
	List() []Tree
}

// TreeBuilder represents a tree builder
type TreeBuilder interface {
	Create() TreeBuilder
	WithEntity(entity entities.Entity) TreeBuilder
	WithGrammar(grammar entities.Identifier) TreeBuilder
	WithLine(line Line) TreeBuilder
	WithSuffix(suffix Trees) TreeBuilder
	Now() (Tree, error)
}

// Tree represents a tree
type Tree interface {
	Entity() entities.Entity
	Grammar() entities.Identifier
	Line() Line
	HasSuffix() bool
	Suffix() Trees
}

// LineBuilder represents the line builder
type LineBuilder interface {
	Create() LineBuilder
	WithEntity(entity entities.Entity) LineBuilder
	WithGrammar(grammar entities.Identifier) LineBuilder
	WithElements(elements Elements) LineBuilder
	Now() (Line, error)
}

// Line represents a line of data
type Line interface {
	Entity() entities.Entity
	Grammar() entities.Identifier
	Elements() Elements
}

// ElementsBuilder represents an elements builder
type ElementsBuilder interface {
	Create() ElementsBuilder
	WithEntity(entity entities.Entity) ElementsBuilder
	WithList(list []Element) ElementsBuilder
	Now() (Elements, error)
}

// Elements represents elements
type Elements interface {
	Entity() entities.Entity
	List() []Element
}

// ElementBuilder represents an element builder
type ElementBuilder interface {
	Create() ElementBuilder
	WithEntity(entity entities.Entity) ElementBuilder
	WithGrammar(grammar entities.Identifier) ElementBuilder
	WithContents(contents Contents) ElementBuilder
	Now() (Element, error)
}

// Element represets an element
type Element interface {
	Entity() entities.Entity
	Grammar() entities.Identifier
	Contents() Contents
}

// ContentsBuilder represents contents builder
type ContentsBuilder interface {
	Create() ContentsBuilder
	WithEntity(entity entities.Entity) ContentsBuilder
	WithList(list []Content) ContentsBuilder
	Now() (Contents, error)
}

// Contents represents contents
type Contents interface {
	Entity() entities.Entity
	List() []Content
}

// ContentBuilder represents a content builder
type ContentBuilder interface {
	Create() ContentBuilder
	WithEntity(entity entities.Entity) ContentBuilder
	WithValue(value Value) ContentBuilder
	WithPrefix(prefix Trees) ContentBuilder
	Now() (Content, error)
}

// Content represents a content
type Content interface {
	Entity() entities.Entity
	Value() Value
	HasPrefix() bool
	Prefix() Trees
}

// ValueBuilder represents a value builder
type ValueBuilder interface {
	Create() ValueBuilder
	WithByte(byte uint8) ValueBuilder
	WithTree(tree Tree) ValueBuilder
	Now() (Value, error)
}

// Value represents value
type Value interface {
	IsByte() bool
	Byte() *uint8
	IsTree() bool
	Tree() Tree
}
