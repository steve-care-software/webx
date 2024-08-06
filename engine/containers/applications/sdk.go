package applications

import (
	"github.com/steve-care-software/webx/engine/containers/domain/containers"
	"github.com/steve-care-software/webx/engine/cursors/domain/hash"
)

// Application represents the container application
type Application interface {
	Begin(name string) (*uint, error)
	Established(identifier uint) bool
	Retrieve(context uint, keyname string) (containers.Container, error)
	Amount(context uint, keyname string) (*uint, error)
	List(context uint, keyname string, index uint64, length uint64) ([]hash.Hash, error)
	Insert(context uint, keyname string, values []hash.Hash) error
	Delete(context uint, keyname string, index uint64, length uint64) error
	Remove(context uint, keyname string) error
	Commit(context uint) error
	DeleteState(context uint, stateIndex uint) error
	RecoverState(context uint, stateIndex uint) error
	StatesAmount(context uint) (*uint, error)
	DeletedStateIndexes(context uint) ([]uint, error)
	Close(context uint) error
	Purge(context uint) error
}
