package vms

import (
	"errors"
	"fmt"

	interpreter_applications "github.com/steve-care-software/webx/interpreters/applications"
	"github.com/steve-care-software/webx/interpreters/domain/results"
	"github.com/steve-care-software/webx/programs/domain/programs/modules"
)

type moduleEngineInterpreter struct {
	interpreterApp interpreter_applications.Application
}

func createModuleInterpreter(
	interpreterApp interpreter_applications.Application,
) *moduleEngineInterpreter {
	out := moduleEngineInterpreter{
		interpreterApp: interpreterApp,
	}

	return &out
}

// Execute executes the module
func (app *moduleEngineInterpreter) Execute() map[uint]modules.ExecuteFn {
	parseThenExecute := app.parseThenExecute()
	resultIsValid := app.resultIsValid()
	resultHasValues := app.resultHasValues()
	resultValues := app.resultValues()
	resultHasRemaining := app.resultHasRemaining()
	resultRemaining := app.resultRemaining()
	return map[uint]modules.ExecuteFn{
		ModuleEngineInterpreterParseThenExecute: parseThenExecute,
		ModuleEngineInterpreterResultIsValid:    resultIsValid,
		ModuleEngineInterpreterResultHasValues:  resultHasValues,
		ModuleEngineInterpreterResultValues:     resultValues,
		ModuleEngineInterpreterHasRemaining:     resultHasRemaining,
		ModuleEngineInterpreterRemaining:        resultRemaining,
	}
}

func (app *moduleEngineInterpreter) parseThenExecute() modules.ExecuteFn {
	return func(input map[uint]interface{}) (interface{}, error) {
		params := []interface{}{}
		if paramsList, ok := input[0].([]interface{}); ok {
			params = paramsList
		}

		if script, ok := input[1].([]byte); ok {
			return app.interpreterApp.ParseThenInterpret(params, script)
		}

		str := fmt.Sprintf("the index 1 was expected to contain a []byte script")
		return nil, errors.New(str)
	}
}

func (app *moduleEngineInterpreter) resultIsValid() modules.ExecuteFn {
	return func(input map[uint]interface{}) (interface{}, error) {
		if res, ok := input[0].(results.Result); ok {
			return res.IsValid(), nil
		}

		str := fmt.Sprintf("the index 0 was expected to contain a Result instance")
		return nil, errors.New(str)
	}
}

func (app *moduleEngineInterpreter) resultHasValues() modules.ExecuteFn {
	return func(input map[uint]interface{}) (interface{}, error) {
		if res, ok := input[0].(results.Result); ok {
			return res.HasValues(), nil
		}

		str := fmt.Sprintf("the index 0 was expected to contain a Result instance")
		return nil, errors.New(str)
	}
}

func (app *moduleEngineInterpreter) resultValues() modules.ExecuteFn {
	return func(input map[uint]interface{}) (interface{}, error) {
		if res, ok := input[0].(results.Result); ok {
			return res.Values(), nil
		}

		str := fmt.Sprintf("the index 0 was expected to contain a Result instance")
		return nil, errors.New(str)
	}
}

func (app *moduleEngineInterpreter) resultHasRemaining() modules.ExecuteFn {
	return func(input map[uint]interface{}) (interface{}, error) {
		if res, ok := input[0].(results.Result); ok {
			return res.HasRemaining(), nil
		}

		str := fmt.Sprintf("the index 0 was expected to contain a Result instance")
		return nil, errors.New(str)
	}
}

func (app *moduleEngineInterpreter) resultRemaining() modules.ExecuteFn {
	return func(input map[uint]interface{}) (interface{}, error) {
		if res, ok := input[0].(results.Result); ok {
			return res.Remaining(), nil
		}

		str := fmt.Sprintf("the index 0 was expected to contain a Result instance")
		return nil, errors.New(str)
	}
}
