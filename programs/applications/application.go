package applications

import (
	"github.com/steve-care-software/webx/blockchains/domain/cryptography/hash"
	"github.com/steve-care-software/webx/programs/domain/compilers"
	"github.com/steve-care-software/webx/programs/domain/programs"
	"github.com/steve-care-software/webx/programs/domain/programs/modules"
)

type application struct {
}

func createApplication() Application {
	out := application{}
	return &out
}

// New creates a new application
func (app *application) New(name string) error {
	return nil
}

// Retrieve retrieves a program by hash
func (app *application) Retrieve(context uint, hash hash.Hash) (programs.Program, error) {
	return nil, nil
}

// Scan scans the database for a program that can receive a given input and returns the requested output
func (app *application) Scan(context uint, input map[string]interface{}, output map[string]interface{}) (programs.Program, error) {
	return nil, nil
}

// Insert inserts a program
func (app *application) Insert(context uint, program programs.Program) error {
	return nil
}

// InsertAll inserts a list of programs
func (app *application) InsertAll(context uint, programs []programs.Program) error {
	return nil
}

// Compile compiles a script to a program
func (app *application) Compile(compiler compilers.Compiler, script []byte) (programs.Program, error) {
	return nil, nil
}

// Execute executes a program
func (app *application) Execute(input map[string]interface{}, modules modules.Modules, program programs.Program) (map[string]interface{}, error) {
	return nil, nil
}
