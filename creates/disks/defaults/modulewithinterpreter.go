package defaults

import (
	"errors"
	"fmt"

	"github.com/steve-care-software/webx/applications"
	creates_module "github.com/steve-care-software/webx/applications/creates/modules"
	"github.com/steve-care-software/webx/domain/programs/modules"
)

type moduleWithInterpreter struct {
	application       applications.Application
	builder           modules.Builder
	moduleBuilder     modules.ModuleBuilder
	additionalModules modules.Modules
}

func createModuleWithInterpreter(
	application applications.Application,
	builder modules.Builder,
	moduleBuilder modules.ModuleBuilder,
	additionalModules modules.Modules,
) creates_module.Application {
	out := moduleWithInterpreter{
		application:       application,
		builder:           builder,
		moduleBuilder:     moduleBuilder,
		additionalModules: additionalModules,
	}

	return &out
}

// Execute executes the application
func (app *moduleWithInterpreter) Execute() (modules.Modules, error) {
	list := []modules.Module{}
	interpreter, err := app.interpreter()
	if err != nil {
		return nil, err
	}

	list = append(list, app.additionalModules.List()...)
	list = append(list, interpreter...)
	return app.builder.Create().WithList(list).Now()
}

func (app *moduleWithInterpreter) interpreter() ([]modules.Module, error) {
	parseThenInterpret, err := app.parseThenInterpret()
	if err != nil {
		return nil, err
	}

	return []modules.Module{
		parseThenInterpret,
	}, nil
}

func (app *moduleWithInterpreter) parseThenInterpret() (modules.Module, error) {
	name := "parseThenInterpret"
	fn := func(input map[string]interface{}) (interface{}, error) {
		parameters := map[string]interface{}{}
		if params, ok := input["params"].(map[string]interface{}); ok {
			parameters = params
		}

		if script, ok := input["script"].([]byte); ok {
			output, remaining, err := app.application.ParseThenInterpret(parameters, script)
			if err != nil {
				return nil, err
			}

			if len(remaining) > 0 {
				str := fmt.Sprintf("the remaining script (%s) was NOT expected", remaining)
				return nil, errors.New(str)
			}

			return output, nil
		}

		return nil, errors.New("the script is mandatory in order to execute the interpreter")
	}

	return app.module(name, fn)
}

func (app *moduleWithInterpreter) module(name string, fn modules.ExecuteFn) (modules.Module, error) {
	return app.moduleBuilder.Create().WithName(name).WithFunc(fn).Now()
}
