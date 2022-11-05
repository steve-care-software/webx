package elements

import (
	"github.com/steve-care-software/webx/domain/databases/entities"
)

// Builder represents an elemnt builder
type Builder interface {
	Create() Builder
	WithEntity(entity entities.Entity) Builder
	WithCardinality(cardinality entities.Identifier) Builder
	WithValue(value uint8) Builder
	WithExternal(external entities.Identifier) Builder
	WithToken(token entities.Identifier) Builder
	WithEverything(everything entities.Identifier) Builder
	WithRecursive(recursive entities.Identifier) Builder
	Now() (Element, error)
}

// Element represents an element
type Element interface {
	Entity() entities.Entity
	Cardinality() entities.Identifier
	Content() Content
}

// Content represents an element content
type Content interface {
	IsValue() bool
	Value() *uint8
	IsExternal() bool
	External() entities.Identifier
	IsToken() bool
	Token() entities.Identifier
	IsEverything() bool
	Everything() entities.Identifier
	IsRecursive() bool
	Recursive() entities.Identifier
}
