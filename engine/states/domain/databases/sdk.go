package databases

import (
	"github.com/steve-care-software/webx/engine/states/domain/databases/commits"
	"github.com/steve-care-software/webx/engine/states/domain/databases/metadatas"
	"github.com/steve-care-software/webx/engine/states/domain/databases/pointers"
	"github.com/steve-care-software/webx/engine/states/domain/files"
	"github.com/steve-care-software/webx/engine/states/domain/hash"
)

// NewRepository creates a new repository
func NewRepository(
	fileRepository files.Repository,
	commitRepository commits.Repository,
	pointerAdapter pointers.Adapter,
) Repository {
	databaseBuilder := NewBuilder()
	return createRepository(
		fileRepository,
		commitRepository,
		pointerAdapter,
		databaseBuilder,
	)
}

// NewService creates a new service
func NewService(
	repository Repository,
	fileService files.Service,
	commitService commits.Service,
	pointerAdapter pointers.Adapter,
) Service {
	pointerBuilder := pointers.NewBuilder()
	return createService(
		repository,
		fileService,
		commitService,
		pointerAdapter,
		pointerBuilder,
	)
}

// NewBuilder creates a new database builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// Builder represents a database builder
type Builder interface {
	Create() Builder
	WithHead(head commits.Commit) Builder
	WithMetaData(metaData metadatas.MetaData) Builder
	Now() (Database, error)
}

// Database represents a database
type Database interface {
	Hash() hash.Hash
	Head() commits.Commit
	MetaData() metadatas.MetaData
}

// Repository represents a database repository
type Repository interface {
	Exists(path []string) bool
	Retrieve(path []string) (Database, error)
}

// Service represents a database service
type Service interface {
	Begin(path []string) error
	Save(database Database) error
	SaveAll(list []Database) error
	End(path []string) error
}
