package grammars

import (
	"github.com/steve-care-software/webx/domain/databases/entities"
)

// Builder represents a grammar
type Builder interface {
	Create() Builder
	WithEntity(entity entities.Entity) Builder
	WithRoot(root entities.Identifier) Builder
	WithChannels(channels []entities.Identifier) Builder
	Now() (Grammar, error)
}

// Grammar represents a grammar
type Grammar interface {
	Entity() entities.Entity
	Root() entities.Identifier
	HasChannels() bool
	Channels() []entities.Identifier
}
