package applications

import (
	"github.com/steve-care-software/webx/engine/cursors/domain/hash"
	"github.com/steve-care-software/webx/engine/vms/domain/instances/executions"
)

// RemoteBuilder represents a remote application builder
type RemoteBuilder interface {
	Create() RemoteBuilder
	WithHost(host string) RemoteBuilder
	Now() (Application, error)
}

// LocalBuilder represents a local application builder
type LocalBuilder interface {
	Create() LocalBuilder
	WithBasePath(basePath []string) LocalBuilder
	Now() (Application, error)
}

// Application represents an application
type Application interface {
	Begin(keyname string) (hash.Hash, error)
	Execute(identifier hash.Hash, input []byte) ([]byte, error)
	Commit(identifier hash.Hash) error
	Executions(identifier hash.Hash) ([]executions.Executions, error)
	Sessions() ([]hash.Hash, error)
	Delete(identifier hash.Hash) error
	DeleteState(identifier hash.Hash, stateIndex uint) error
	RecoverState(identifier hash.Hash, stateIndex uint) error
	StatesAmount(identifier hash.Hash) (*uint, error)
	DeletedStateIndexes(identifier hash.Hash) ([]uint, error)
	Close(identifier hash.Hash) error
	Purge(identifier hash.Hash) error
}
