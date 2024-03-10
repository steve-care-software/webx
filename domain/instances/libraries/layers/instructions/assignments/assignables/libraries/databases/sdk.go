package databases

import (
	"github.com/steve-care-software/datastencil/domain/instances/libraries/layers/instructions/assignments/assignables/libraries/databases/repositories"
	"github.com/steve-care-software/datastencil/domain/instances/libraries/layers/instructions/assignments/assignables/libraries/databases/services"
)

// Builder represents a database builder
type Builder interface {
	Create() Builder
	WithRepository(repository repositories.Repository) Builder
	WithService(service services.Service) Builder
	Now() (Database, error)
}

// Database represents a database instruction
type Database interface {
	IsRepository() bool
	Repository() repositories.Repository
	IsService() bool
	Service() services.Service
}
