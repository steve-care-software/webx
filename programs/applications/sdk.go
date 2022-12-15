package applications

import (
	"github.com/steve-care-software/webx/blockchains/domain/cryptography/hash"
	grammars_application "github.com/steve-care-software/webx/roots/applications/grammars"
	"github.com/steve-care-software/webx/programs/domain/programs"
	"github.com/steve-care-software/webx/programs/domain/programs/modules"
)

const (
	// KindProgram represents the program kind
	KindProgram = grammars_application.KindGrammar + 1

	// KindApplication represents the application kind
	KindApplication

	// KindInstruction represents the instruction kind
	KindInstruction

	// KindValue represents the value kind
	KindValue
)

// ScanCallbackFn represents the scan callback func
type ScanCallbackFn func(output map[string]interface{}) bool

// Application represents a program application
type Application interface {
	Database
	Software
}

// Software represents the program software application
type Software interface {
	Execute(input map[uint]interface{}, program programs.Program) (map[uint]interface{}, error)
}

// Database represents the program database application
type Database interface {
	Retrieve(context uint, hash hash.Hash, modules modules.Modules) (programs.Program, error)
	Scan(context uint, input map[string]interface{}, callbackFn ScanCallbackFn) (programs.Program, error)
	Insert(context uint, program programs.Program) error
}
