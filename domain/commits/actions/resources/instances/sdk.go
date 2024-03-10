package instances

import (
	"github.com/steve-care-software/datastencil/domain/accounts/signers"
	"github.com/steve-care-software/datastencil/domain/commits/actions/resources/instances/skeletons"
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/libraries/layers/instructions/assignments/assignables/queries"
)

// Adapter represents the instance adapter
type Adapter interface {
	ToData(ins Instance) ([]byte, error)
	ToInstance(data []byte) (Instance, error)
}

// Instance represents an instance
type Instance interface {
	Hash() hash.Hash
}

// RepositoryBuilder represents the repository builder
type RepositoryBuilder interface {
	Create() RepositoryBuilder
	WithSkeleton(skeleton skeletons.Skeleton) RepositoryBuilder
	WithSigner(signer signers.Signer) RepositoryBuilder
	Now() (Repository, error)
}

// Repository represents an instance repository
type Repository interface {
	// Height returns the current commit height
	Height() (uint, error)

	// List returns the hashes list related to the query
	List(query queries.Query) ([]hash.Hash, error)

	// Exists returns true if the instance exists, false otherwise
	Exists(path []string, hash hash.Hash) (bool, error)

	// Retrieve returns the instance by query
	Retrieve(query queries.Query) (Instance, error)

	// RetrieveByPathAndHash returns the instance by path and hash
	RetrieveByPathAndHash(path []string, hash hash.Hash) (Instance, error)
}

// ServiceBuilder represents the service builder
type ServiceBuilder interface {
	Create() ServiceBuilder
	WithSkeleton(skeleton skeletons.Skeleton) ServiceBuilder
	WithSigner(signer signers.Signer) ServiceBuilder
	Now() (Service, error)
}

// Service represents a an instance service
type Service interface {
	// Init initializes the service
	Init() error

	// Begin begins a transaction
	Begin() (uint, error)

	// Insert inserts an instance
	Insert(context uint, ins Instance, path []string) error

	// Delete deletes an instance
	Delete(context uint, path []string, hash hash.Hash) error

	// Commit commits actions
	Commit(context uint) error

	// Cancel cancels a context
	Cancel(context uint) error

	// Revert reverts the state of the last commit
	Revert() error

	// Revert reverts the state of the commit to the provided index
	RevertToIndex(toIndex uint) error
}
