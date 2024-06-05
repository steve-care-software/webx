package mocks

import (
	"github.com/steve-care-software/datastencil/domain/encryptors"
	"github.com/steve-care-software/datastencil/domain/instances"
	"github.com/steve-care-software/datastencil/domain/instances/databases"
)

// NewEncryptor creates a new encryptor
func NewEncryptor(
	encrypt map[string]map[string][]byte,
	decrypt map[string]map[string][]byte,
) encryptors.Encryptor {
	return createEncryptor(
		encrypt,
		decrypt,
	)
}

// NewInstanceAdapter creates a new instance adapter
func NewInstanceAdapter(
	toData map[string][]byte,
	toInstance map[string]instances.Instance,
) instances.Adapter {
	return createInstanceAdapter(
		toData,
		toInstance,
	)
}

// NewDatabaseRepository creates a new database repository
func NewDatabaseRepository(
	list [][]string,
	database databases.Database,
	errorIns error,
) databases.Repository {
	return createDatabaseRepository(
		list,
		database,
		errorIns,
	)
}
