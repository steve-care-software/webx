package elements

import "github.com/steve-care-software/webx/domain/databases/entities"

// Builder represents an elements builder
type Builder interface {
	Create() Builder
	WithEntity(entity entities.Entity) Builder
	WithList(list []Element) Builder
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
	WithContents(contents entities.Identifiers) ElementBuilder
	Now() (Element, error)
}

// Element represets an element
type Element interface {
	Entity() entities.Entity
	Grammar() entities.Identifier
	Contents() entities.Identifiers
}
