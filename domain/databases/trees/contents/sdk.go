package contents

import "github.com/steve-care-software/webx/domain/databases/entities"

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents a content builder
type Builder interface {
	Create() Builder
	WithEntity(entity entities.Entity) Builder
	WithValue(value Value) Builder
	WithPrefix(prefix entities.Identifiers) Builder
	Now() (Content, error)
}

// Content represents a content
type Content interface {
	Entity() entities.Entity
	Value() Value
	HasPrefix() bool
	Prefix() entities.Identifiers
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
