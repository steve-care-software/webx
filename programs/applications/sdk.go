package applications

import (
	"github.com/steve-care-software/webx/blockchains/applications"
	"github.com/steve-care-software/webx/blockchains/domain/cryptography/hash"
	"github.com/steve-care-software/webx/programs/domain/instructions"
	"github.com/steve-care-software/webx/programs/domain/programs"
)

// Application represents a program application
type Application interface {
	New(name string) error
	applications.Database
	Database
	Software
}

// Software represents the program software application
type Software interface {
	Compile(instructions instructions.Instructions) (programs.Program, error)
	Execute(input map[string]interface{}, program programs.Program) (map[string]interface{}, error)
}

// Database represents the program database application
type Database interface {
	Retrieve(context uint, hash hash.Hash) (programs.Program, error)
	Search(context uint, input map[string]interface{}, output map[string]interface{}) (programs.Program, error)
	Insert(context uint, program programs.Program) error
	InsertAll(context uint, programs []programs.Program) error
}
