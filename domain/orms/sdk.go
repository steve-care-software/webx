package orms

import (
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/orms/skeletons"
)

// Instance represents an instance
type Instance interface {
	Hash() hash.Hash
}

// Actions represents actions
type Actions interface {
	HasInsert() bool
	Insert() Resources
	HasDelete() bool
	Delete() Pointers
}

// Resources represents resources
type Resources interface {
	Hash() hash.Hash
	List() []Resource
}

// Resource represents a resource
type Resource interface {
	Path() []string
	Instance() Instance
}

// Pointers represents pointers
type Pointers interface {
	Hash() hash.Hash
	List() []Pointer
}

// Pointer represents a pointer
type Pointer interface {
	Path() []string
	Hash() hash.Hash
}

// RepositoryBuilder represents the repository builder
type RepositoryBuilder interface {
	Create() RepositoryBuilder
	WithSkeleton(skeleton skeletons.Skeleton) RepositoryBuilder
	Now() (Repository, error)
}

// Repository represents an instance repository
type Repository interface {
	// Retrieve retrieves an instance by path and hash
	Retrieve(path []string, hash hash.Hash) (Instance, error)

	// List retrieves a list of to hashes
	List(fromPath []string, toPath []string, fromHash hash.Hash) ([]hash.Hash, error)
}

// ServiceBuilder represents the service builder
type ServiceBuilder interface {
	Create() ServiceBuilder
	WithSkeleton(skeleton skeletons.Skeleton) ServiceBuilder
	Now() (Service, error)
}

// Service represents a an instance service
type Service interface {
	// Init initializes the service
	Init() error

	// Insert inserts an instance
	Insert(ins Instance, path []string) error

	// Delete deletes an instance
	Delete(path []string, hash hash.Hash) error
}
