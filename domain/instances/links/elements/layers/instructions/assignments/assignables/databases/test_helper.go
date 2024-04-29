package databases

import (
	"github.com/steve-care-software/datastencil/domain/instances/links/elements/layers/instructions/assignments/assignables/databases/repositories"
	"github.com/steve-care-software/datastencil/domain/instances/links/elements/layers/instructions/assignments/assignables/databases/services"
)

// NewDatabaseWithRepositoryForTests creates a new database with repository for tests
func NewDatabaseWithRepositoryForTests(repository repositories.Repository) Database {
	ins, err := NewBuilder().Create().WithRepository(repository).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewDatabaseWithServiceForTests creates a new database with service for tests
func NewDatabaseWithServiceForTests(service services.Service) Database {
	ins, err := NewBuilder().Create().WithService(service).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
