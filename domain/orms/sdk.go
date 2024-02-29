package orms

import (
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/orms/skeletons"
)

// Instance represents an instance
type Instance interface {
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
