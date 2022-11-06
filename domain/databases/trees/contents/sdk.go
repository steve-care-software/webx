package contents

import "github.com/steve-care-software/webx/domain/databases/entities"

// Builder represents contents builder
type Builder interface {
	Create() Builder
	WithEntity(entity entities.Entity) Builder
	WithList(list []Content) Builder
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
	WithPrefix(prefix entities.Identifier) ContentBuilder
	Now() (Content, error)
}

// Content represents a content
type Content interface {
	Entity() entities.Entity
	Value() Value
	HasPrefix() bool
	Prefix() entities.Identifier
}

// ValueBuilder represents a value builder
type ValueBuilder interface {
	Create() ValueBuilder
	WithByte(byte byte) ValueBuilder
	WithTree(tree entities.Identifier) ValueBuilder
	Now() (Value, error)
}

// Value represents value
type Value interface {
	IsByte() bool
	Byte() *byte
	IsTree() bool
	Tree() entities.Identifier
}
