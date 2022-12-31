package compilers

import (
	"github.com/steve-care-software/webx/compilers/domain/compilers"
	"github.com/steve-care-software/webx/databases/applications"
	"github.com/steve-care-software/webx/databases/domain/cryptography/hash"
)

type application struct {
	databaseApp applications.Application
}

func createApplication(
	databaseApp applications.Application,
) Application {
	out := application{
		databaseApp: databaseApp,
	}

	return &out
}

// List lists the compilers
func (app *application) List(ontext uint) ([]hash.Hash, error) {
	return nil, nil
}

// Retrieve retrieves a compiler by hash
func (app *application) Retrieve(context uint, hash hash.Hash) (compilers.Compiler, error) {
	return nil, nil
}

// Insert inserts a compiler
func (app *application) Insert(context uint, compiler compilers.Compiler) error {
	return nil
}

// InsertAll inserts a list of compilers
func (app *application) InsertAll(context uint, compilers []compilers.Compiler) error {
	return nil
}
