package applications

import (
	"github.com/steve-care-software/webx/engine/cursors/domain/hash"
	"github.com/steve-care-software/webx/engine/entities/domain/entities"
	hash_applications "github.com/steve-care-software/webx/engine/hashes/applications"
)

// NewBuilder creates a new builder
func NewBuilder(
	entityAdapter entities.Adapter,
) Builder {
	return createBuilder(
		entityAdapter,
	)
}

// Builder represents an application builder
type Builder interface {
	Create() Builder
	WithHash(hashApp hash_applications.Application) Builder
	Now() (Application, error)
}

// Application represents the database application
type Application interface {
	Begin(name string) (*uint, error)
	Established(identifier uint) bool
	Retrieve(context uint, hash hash.Hash) (entities.Entity, error)
	Insert(context uint, entity entities.Entity) error
	Delete(context uint, hash hash.Hash) error
	Commit(context uint) error
	DeleteState(context uint, stateIndex uint) error
	RecoverState(context uint, stateIndex uint) error
	StatesAmount(context uint) (*uint, error)
	DeletedStateIndexes(context uint) ([]uint, error)
	Close(context uint) error
	Purge(context uint) error
}
