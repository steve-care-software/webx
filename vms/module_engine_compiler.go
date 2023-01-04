package vms

import (
	compiler_applications "github.com/steve-care-software/webx/compilers/applications"
	"github.com/steve-care-software/webx/programs/domain/programs/modules"
)

type moduleEngineCompiler struct {
	compilerApp compiler_applications.Application
}

func createModuleCompiler(
	compilerApp compiler_applications.Application,
) *moduleEngineCompiler {
	out := moduleEngineCompiler{
		compilerApp: compilerApp,
	}

	return &out
}

// Execute executes the module
func (app *moduleEngineCompiler) Execute() map[uint]modules.ExecuteFn {
	execute := app.execute()
	return map[uint]modules.ExecuteFn{
		ModuleEngineCompilerExecute: execute,
	}
}

func (app *moduleEngineCompiler) execute() modules.ExecuteFn {
	return func(input map[uint]interface{}) (interface{}, error) {
		/*if paramsList, ok := input[0].([]interface{}); ok {
			if script, ok := input[1].([]byte); ok {
				return app.compilerApp.ParseThenInterpret(paramsList, script)
			}

			str := fmt.Sprintf("the index 1 was expected to contain a []byte script")
			return nil, errors.New(str)
		}

		str := fmt.Sprintf("the index 0 was expected to be valid and contain a list")
		return nil, errors.New(str)*/
		return nil, nil
	}
}
