package defaults

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	creates_module "github.com/steve-care-software/syntax/applications/engines/creates/modules"
	"github.com/steve-care-software/syntax/domain/syntax/grammars/values"
	"github.com/steve-care-software/syntax/domain/syntax/programs/applications/modules"
)

type module struct {
	builder             modules.Builder
	moduleBuilder       modules.ModuleBuilder
	grammarValueBuilder values.Builder
}

func createModule(
	builder modules.Builder,
	moduleBuilder modules.ModuleBuilder,
	grammarValueBuilder values.Builder,
) creates_module.Application {
	out := module{
		builder:             builder,
		moduleBuilder:       moduleBuilder,
		grammarValueBuilder: grammarValueBuilder,
	}

	return &out
}

// Execute executes the application
func (app *module) Execute() (modules.Modules, error) {
	list := []modules.Module{}
	engine, err := app.engine()
	if err != nil {
		return nil, err
	}

	list = append(list, engine...)
	return app.builder.Create().WithList(list).Now()
}

func (app *module) engine() ([]modules.Module, error) {
	list := []modules.Module{}
	grammar, err := app.grammar()
	if err != nil {
		return nil, err
	}

	list = append(list, grammar...)
	return list, nil
}

func (app *module) grammar() ([]modules.Module, error) {
	value, err := app.grammarValue()
	if err != nil {
		return nil, err
	}

	return []modules.Module{
		value,
	}, nil
}

func (app *module) grammarValue() (modules.Module, error) {
	name := "engineGrammarValue"
	fn := func(input map[string]interface{}) (interface{}, error) {
		if name, ok := input["name"].(string); ok {
			if number, ok := input["number"].(string); ok {
				intNumber, err := strconv.Atoi(strings.TrimSpace(number))
				if err != nil {
					return nil, err
				}

				if intNumber < 0 {
					return nil, errors.New("the number cannot be smaller than 0")
				}

				if intNumber > 255 {
					return nil, errors.New("the number cannot be bigger than 255")
				}

				return app.grammarValueBuilder.Create().
					WithName(name).
					WithNumber(byte(intNumber)).
					Now()
			}

			str := fmt.Sprintf("the number was expected to be valid and contain a string")
			return nil, errors.New(str)

		}

		str := fmt.Sprintf("the name was expected to be valid and contain a string")
		return nil, errors.New(str)
	}

	return app.module(name, fn)
}

func (app *module) module(name string, fn modules.ExecuteFn) (modules.Module, error) {
	return app.moduleBuilder.Create().WithName(name).WithFunc(fn).Now()
}
