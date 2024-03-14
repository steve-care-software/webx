package libraries

import (
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/libraries/links/layers/instructions/assignments/assignables/libraries/compilers"
	"github.com/steve-care-software/datastencil/domain/instances/libraries/links/layers/instructions/assignments/assignables/libraries/databases"
)

type library struct {
	hash     hash.Hash
	compiler compilers.Compiler
	database databases.Database
}

func createLibraryWithCompiler(
	hash hash.Hash,
	compiler compilers.Compiler,
) Library {
	return createLibraryInternally(hash, compiler, nil)
}

func createLibraryWithDatabase(
	hash hash.Hash,
	database databases.Database,
) Library {
	return createLibraryInternally(hash, nil, database)
}

func createLibraryInternally(
	hash hash.Hash,
	compiler compilers.Compiler,
	database databases.Database,
) Library {
	out := library{
		hash:     hash,
		compiler: compiler,
		database: database,
	}

	return &out
}

// Hash returns the hash
func (obj *library) Hash() hash.Hash {
	return obj.hash
}

// IsCompiler returns true if there is a compiler, false otherwise
func (obj *library) IsCompiler() bool {
	return obj.compiler != nil
}

// Compiler returns the compiler, if any
func (obj *library) Compiler() compilers.Compiler {
	return obj.compiler
}

// IsDatabase returns true if there is a database, false otherwise
func (obj *library) IsDatabase() bool {
	return obj.database != nil
}

// Database returns the database, if any
func (obj *library) Database() databases.Database {
	return obj.database
}
