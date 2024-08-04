package applications

import (
	bytes_applications "github.com/steve-care-software/webx/engine/bytes/applications"
	"github.com/steve-care-software/webx/engine/bytes/domain/states/pointers/delimiters"
	"github.com/steve-care-software/webx/engine/hashes/domain/hash"
	"github.com/steve-care-software/webx/engine/hashes/domain/pointers"
)

const contextIdentifierUndefinedPattern = "the context identifier (%d) does not exists"

// NewBuilder creates a new application builder
func NewBuilder(
	pointerAdapter pointers.Adapter,
) Builder {
	pointersBuilder := pointers.NewBuilder()
	pointerBuilder := pointers.NewPointerBuilder()
	delimiterBuilder := delimiters.NewDelimiterBuilder()
	return createBuilder(
		pointerAdapter,
		pointersBuilder,
		pointerBuilder,
		delimiterBuilder,
	)
}

// Builder represents an application builder
type Builder interface {
	Create() Builder
	WithBytes(bytesApp bytes_applications.Application) Builder
	Now() (Application, error)
}

// Application represents an hash application
type Application interface {
	Begin(name string) (*uint, error)
	Established(context uint) bool
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
