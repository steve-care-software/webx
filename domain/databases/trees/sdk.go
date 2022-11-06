package trees

import "github.com/steve-care-software/webx/domain/databases/entities"

// Builder represents a tree builder
type Builder interface {
	Create() Builder
	WithEntity(entity entities.Entity) Builder
	WithGrammar(grammar entities.Identifier) Builder
	WithLine(line entities.Identifier) Builder
	WithSuffix(suffix entities.Identifiers) Builder
	Now() (Tree, error)
}

// Tree represents a tree
type Tree interface {
	Entity() entities.Entity
	Grammar() entities.Identifier
	Line() entities.Identifier
	HasSuffix() bool
	Suffix() entities.Identifiers
}
