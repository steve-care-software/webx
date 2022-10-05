package defaults

import (
	creates_module "github.com/steve-care-software/syntax/applications/engines/creates/modules"
	"github.com/steve-care-software/syntax/domain/syntax/programs/instructions/applications/modules"
)

type module struct {
}

func createModule() creates_module.Application {
	out := module{}
	return &out
}

// Execute executes the application
func (app *module) Execute() (modules.Modules, error) {
	return nil, nil
}
