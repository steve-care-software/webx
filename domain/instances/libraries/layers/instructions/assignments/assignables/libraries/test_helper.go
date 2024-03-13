package libraries

import (
	"github.com/steve-care-software/datastencil/domain/instances/libraries/layers/instructions/assignments/assignables/libraries/compilers"
	"github.com/steve-care-software/datastencil/domain/instances/libraries/layers/instructions/assignments/assignables/libraries/databases"
)

// NewLibraryWithCompilerForTests creates a new library with compiler for tests
func NewLibraryWithCompilerForTests(
	compiler compilers.Compiler,
) Library {
	ins, err := NewBuilder().Create().WithCompiler(compiler).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewLibraryWithDatabaseForTests creates a new library with database for tests
func NewLibraryWithDatabaseForTests(
	database databases.Database,
) Library {
	ins, err := NewBuilder().Create().WithDatabase(database).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
