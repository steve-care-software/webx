package commits

import (
	"github.com/steve-care-software/webx/engine/databases/entities/domain/hash"
	"github.com/steve-care-software/webx/engine/states/domain/databases/commits/executions"
	"github.com/steve-care-software/webx/engine/states/domain/files"
)

// NewRepository creates a new repository
func NewRepository(
	adapter Adapter,
	fileRepository files.Repository,
) Repository {
	return createRepository(
		adapter,
		fileRepository,
	)
}

// NewService creates a new service
func NewService(
	adapter Adapter,
	fileService files.Service,
) Service {
	return createService(
		adapter,
		fileService,
	)
}

// NewBuilder creates a new commit builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// Adapter represents a commit adapter
type Adapter interface {
	ToBytes(ins Commit) ([]byte, error)
	ToInstance(bytes []byte) (Commit, error)
}

// Builder represents a commit builder
type Builder interface {
	Create() Builder
	WithExecutions(executions executions.Executions) Builder
	WithParent(parent hash.Hash) Builder
	Now() (Commit, error)
}

// Commit represents a commit
type Commit interface {
	Hash() hash.Hash
	Executions() executions.Executions
	HasParent() bool
	Parent() hash.Hash
}

// RepositoryBuilder represents a repository builder
type RepositoryBuilder interface {
	Create() RepositoryBuilder
	WithBasePath(basePath []string) RepositoryBuilder
	Now() (Repository, error)
}

// Repository represents a commit repository
type Repository interface {
	Retrieve(hash hash.Hash) (Commit, error)
}

// ServiceBuilder represents a service builder
type ServiceBuilder interface {
	Create() ServiceBuilder
	WithBasePath(basePath []string) ServiceBuilder
	Now() (Service, error)
}

// Service represents the commit service
type Service interface {
	Save(ins Commit) error
}
