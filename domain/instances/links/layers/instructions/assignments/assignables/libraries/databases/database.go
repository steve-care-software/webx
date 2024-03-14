package databases

import (
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/links/layers/instructions/assignments/assignables/libraries/databases/repositories"
	"github.com/steve-care-software/datastencil/domain/instances/links/layers/instructions/assignments/assignables/libraries/databases/services"
)

type database struct {
	hash       hash.Hash
	repository repositories.Repository
	service    services.Service
}

func createDatabaseWithRepository(
	hash hash.Hash,
	repository repositories.Repository,
) Database {
	return createDatabaseInternally(hash, repository, nil)
}

func createDatabaseWithService(
	hash hash.Hash,
	service services.Service,
) Database {
	return createDatabaseInternally(hash, nil, service)
}

func createDatabaseInternally(
	hash hash.Hash,
	repository repositories.Repository,
	service services.Service,
) Database {
	out := database{
		hash:       hash,
		repository: repository,
		service:    service,
	}

	return &out
}

// Hash returns the hash
func (obj *database) Hash() hash.Hash {
	return obj.hash
}

// IsRepository returns true if there is a repository, false otherwise
func (obj *database) IsRepository() bool {
	return obj.repository != nil
}

// Repository returns the repository, if any
func (obj *database) Repository() repositories.Repository {
	return obj.repository
}

// IsService returns true if there is a service, false otherwise
func (obj *database) IsService() bool {
	return obj.service != nil
}

// Service returns the service, if any
func (obj *database) Service() services.Service {
	return obj.service
}
