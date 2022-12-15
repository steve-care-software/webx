package elements

import "github.com/steve-care-software/webx/domain/databases/entities"

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents an element builder
type Builder interface {
	Create() Builder
	WithEntity(entity entities.Entity) Builder
	WithGrammar(grammar entities.Identifier) Builder
	WithContents(contents entities.Identifiers) Builder
	Now() (Element, error)
}

// Element represets an element
type Element interface {
	Entity() entities.Entity
	Grammar() entities.Identifier
	Contents() entities.Identifiers
}
