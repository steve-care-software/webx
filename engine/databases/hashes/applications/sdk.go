package applications

import "github.com/steve-care-software/webx/engine/databases/entities/domain/hash"

// Application represents an hash application
type Application interface {
	Begin(name string) (*uint, error)
	Retrieve(context uint, hash hash.Hash) ([]byte, error)
	RetrieveAll(context uint, hashes []hash.Hash) ([][]byte, error)
	Insert(context uint, data []byte) error
	InsertAll(context uint, data [][]byte) error
	Delete(context uint, hash hash.Hash) error
	DeleteAll(context uint, hashes []hash.Hash) error
	Commit(context uint) error
	DeleteState(context uint, stateIndex uint) error
	RecoverState(context uint, stateIndex uint) error
	StatesAmount(context uint) (*uint, error)
	DeletedStateIndexes(context uint) ([]uint, error)
	Close(context uint) error
	Purge(context uint) error
}
