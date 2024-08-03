package applications

import (
	bytes_applications "github.com/steve-care-software/webx/engine/databases/bytes/applications"
	"github.com/steve-care-software/webx/engine/databases/bytes/domain/states/pointers/delimiters"
	"github.com/steve-care-software/webx/engine/databases/hashes/domain/hash"
	"github.com/steve-care-software/webx/engine/databases/hashes/domain/pointers"
)

const contextIdentifierUndefinedPattern = "the context identifier (%d) does not exists"

// NewApplication creates a new application
func NewApplication(
	bytesApp bytes_applications.Application,
	pointerAdapter pointers.Adapter,
) Application {
	pointersBuilder := pointers.NewBuilder()
	pointerBuilder := pointers.NewPointerBuilder()
	delimiterBuilder := delimiters.NewDelimiterBuilder()
	return createApplication(
		bytesApp,
		pointerAdapter,
		pointersBuilder,
		pointerBuilder,
		delimiterBuilder,
	)
}

// Application represents an hash application
type Application interface {
	Begin(name string) (*uint, error)
	Retrieve(context uint, hash hash.Hash) ([]byte, error)
	Insert(context uint, hash hash.Hash, data []byte) error
	Delete(context uint, hash hash.Hash) error
	Commit(context uint) error
	DeleteState(context uint, stateIndex uint) error
	RecoverState(context uint, stateIndex uint) error
	StatesAmount(context uint) (*uint, error)
	DeletedStateIndexes(context uint) ([]uint, error)
	Close(context uint) error
	Purge(context uint) error
}
