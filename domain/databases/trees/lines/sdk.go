package lines

import "github.com/steve-care-software/webx/domain/databases/entities"

// Builder represents the line builder
type Builder interface {
	Create() Builder
	WithEntity(entity entities.Entity) Builder
	WithGrammar(grammar entities.Identifier) Builder
	WithElements(elements entities.Identifiers) Builder
	Now() (Line, error)
}

// Line represents a line of data
type Line interface {
	Entity() entities.Entity
	Grammar() entities.Identifier
	Elements() entities.Identifiers
}
