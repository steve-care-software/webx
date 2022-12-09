package applications

import (
	"github.com/steve-care-software/webx/blockchains/domain/cryptography/hash"
	"github.com/steve-care-software/webx/programs/domain/programs"
	"github.com/steve-care-software/webx/programs/domain/programs/modules"
)

// ScanCallbackFn represents the scan callback func
type ScanCallbackFn func(output map[string]interface{}) bool

// Application represents a program application
type Application interface {
	New(name string) error
	Database
	Software
}

// Software represents the program software application
type Software interface {
	Execute(input map[string]interface{}, program programs.Program) (map[string]interface{}, error)
}

// Database represents the program database application
type Database interface {
	Retrieve(context uint, hash hash.Hash, modules modules.Modules) (programs.Program, error)
	Scan(context uint, input map[string]interface{}, callbackFn ScanCallbackFn) (programs.Program, error)
	Insert(context uint, program programs.Program) error
}
