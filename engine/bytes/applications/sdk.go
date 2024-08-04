package applications

import (
	"github.com/steve-care-software/webx/engine/bytes/domain/delimiters"
	"github.com/steve-care-software/webx/engine/bytes/domain/states"
)

// Builder represents the application builder
type Builder interface {
	Create() Builder
	WithBasePath(basePath []string) Builder
	Now() (Application, error)
}

// Application represents the database application
type Application interface {
	Begin(name string) (*uint, error)
	Retrieve(context uint, retrival delimiters.Delimiter) ([]byte, error)
	RetrieveAll(context uint, retrievals delimiters.Delimiters) ([][]byte, error)
	Insert(context uint, data []byte) (delimiters.Delimiter, error)
	Delete(context uint, delete delimiters.Delimiter) error
	DeleteAll(context uint, deletes delimiters.Delimiters) error
	Commit(context uint) error
	CommitWithRoot(context uint, root delimiters.Delimiter) error
	DeleteState(context uint, stateIndex uint) error
	RecoverState(context uint, stateIndex uint) error
	StatesAmount(context uint) (*uint, error)
	DeletedStateIndexes(context uint) ([]uint, error)
	Close(context uint) error
	Purge(context uint) error
	States(context uint) (states.States, error)
	Switch(context uint, branchName string) error
	Dive(context uint, childrenBranchName string) error
	Climb(context uint) error
	Merge(context uint) error
	BranchNames(context uint) ([]string, error)
	DeleteBranch(context uint) error
	RecoverBranch(context uint, name string) error
	DeletedBranchNames(context uint) ([]string, error)
	LayerAmount(context uint) (*uint, error)
	DeleteLayer(context uint, layerIndex uint) error
	RecoverLayer(context uint, layerIndex uint) error
	DeletedLayerIndexes(context uint) ([]uint, error)
}
