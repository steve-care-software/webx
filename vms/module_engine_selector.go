package vms

import (
	program_applications "github.com/steve-care-software/webx/programs/applications"
	"github.com/steve-care-software/webx/programs/domain/programs/modules"
	selector_applications "github.com/steve-care-software/webx/selectors/applications"
)

type moduleEngineSelector struct {
	programApp  program_applications.Application
	selectorApp selector_applications.Application
}

func createModuleSelector(
	programApp program_applications.Application,
	selectorApp selector_applications.Application,
) *moduleEngineSelector {
	out := moduleEngineSelector{
		programApp:  programApp,
		selectorApp: selectorApp,
	}

	return &out
}

// Execute executes the module
func (app *moduleEngineSelector) Execute() map[uint]modules.ExecuteFn {
	fetcher := app.fetcher()
	fetchers := app.fetchers()
	contentFn := app.inside()
	inside := app.inside()
	element := app.element()
	token := app.token()
	selectorFn := app.selectorFn()
	selector := app.selector()
	execute := app.execute()
	return map[uint]modules.ExecuteFn{
		ModuleEngineSelectorFetcher:    fetcher,
		ModuleEngineSelectorFetchers:   fetchers,
		ModuleEngineSelectorContentFn:  contentFn,
		ModuleEngineSelectorInside:     inside,
		ModuleEngineSelectorElement:    element,
		ModuleEngineSelectorToken:      token,
		ModuleEngineSelectorSelectorFn: selectorFn,
		ModuleEngineSelector:           selector,
		ModuleEngineExecute:            execute,
	}
}

func (app *moduleEngineSelector) execute() modules.ExecuteFn {
	return func(input map[uint]interface{}) (interface{}, error) {
		/*if paramsList, ok := input[0].([]interface{}); ok {
			if script, ok := input[1].([]byte); ok {
				return app.selectorApp.ParseThenInterpret(paramsList, script)
			}

			str := fmt.Sprintf("the index 1 was expected to contain a []byte script")
			return nil, errors.New(str)
		}

		str := fmt.Sprintf("the index 0 was expected to be valid and contain a list")
		return nil, errors.New(str)*/
		return nil, nil
	}
}

func (app *moduleEngineSelector) selector() modules.ExecuteFn {
	return func(input map[uint]interface{}) (interface{}, error) {
		/*if paramsList, ok := input[0].([]interface{}); ok {
			if script, ok := input[1].([]byte); ok {
				return app.selectorApp.ParseThenInterpret(paramsList, script)
			}

			str := fmt.Sprintf("the index 1 was expected to contain a []byte script")
			return nil, errors.New(str)
		}

		str := fmt.Sprintf("the index 0 was expected to be valid and contain a list")
		return nil, errors.New(str)*/
		return nil, nil
	}
}

func (app *moduleEngineSelector) selectorFn() modules.ExecuteFn {
	return func(input map[uint]interface{}) (interface{}, error) {
		/*if paramsList, ok := input[0].([]interface{}); ok {
			if script, ok := input[1].([]byte); ok {
				return app.selectorApp.ParseThenInterpret(paramsList, script)
			}

			str := fmt.Sprintf("the index 1 was expected to contain a []byte script")
			return nil, errors.New(str)
		}

		str := fmt.Sprintf("the index 0 was expected to be valid and contain a list")
		return nil, errors.New(str)*/
		return nil, nil
	}
}

func (app *moduleEngineSelector) token() modules.ExecuteFn {
	return func(input map[uint]interface{}) (interface{}, error) {
		/*if paramsList, ok := input[0].([]interface{}); ok {
			if script, ok := input[1].([]byte); ok {
				return app.selectorApp.ParseThenInterpret(paramsList, script)
			}

			str := fmt.Sprintf("the index 1 was expected to contain a []byte script")
			return nil, errors.New(str)
		}

		str := fmt.Sprintf("the index 0 was expected to be valid and contain a list")
		return nil, errors.New(str)*/
		return nil, nil
	}
}

func (app *moduleEngineSelector) element() modules.ExecuteFn {
	return func(input map[uint]interface{}) (interface{}, error) {
		/*if paramsList, ok := input[0].([]interface{}); ok {
			if script, ok := input[1].([]byte); ok {
				return app.selectorApp.ParseThenInterpret(paramsList, script)
			}

			str := fmt.Sprintf("the index 1 was expected to contain a []byte script")
			return nil, errors.New(str)
		}

		str := fmt.Sprintf("the index 0 was expected to be valid and contain a list")
		return nil, errors.New(str)*/
		return nil, nil
	}
}

func (app *moduleEngineSelector) inside() modules.ExecuteFn {
	return func(input map[uint]interface{}) (interface{}, error) {
		/*if paramsList, ok := input[0].([]interface{}); ok {
			if script, ok := input[1].([]byte); ok {
				return app.selectorApp.ParseThenInterpret(paramsList, script)
			}

			str := fmt.Sprintf("the index 1 was expected to contain a []byte script")
			return nil, errors.New(str)
		}

		str := fmt.Sprintf("the index 0 was expected to be valid and contain a list")
		return nil, errors.New(str)*/
		return nil, nil
	}
}

func (app *moduleEngineSelector) fetchers() modules.ExecuteFn {
	return func(input map[uint]interface{}) (interface{}, error) {
		/*if paramsList, ok := input[0].([]interface{}); ok {
			if script, ok := input[1].([]byte); ok {
				return app.selectorApp.ParseThenInterpret(paramsList, script)
			}

			str := fmt.Sprintf("the index 1 was expected to contain a []byte script")
			return nil, errors.New(str)
		}

		str := fmt.Sprintf("the index 0 was expected to be valid and contain a list")
		return nil, errors.New(str)*/
		return nil, nil
	}
}

func (app *moduleEngineSelector) fetcher() modules.ExecuteFn {
	return func(input map[uint]interface{}) (interface{}, error) {
		/*if paramsList, ok := input[0].([]interface{}); ok {
			if script, ok := input[1].([]byte); ok {
				return app.selectorApp.ParseThenInterpret(paramsList, script)
			}

			str := fmt.Sprintf("the index 1 was expected to contain a []byte script")
			return nil, errors.New(str)
		}

		str := fmt.Sprintf("the index 0 was expected to be valid and contain a list")
		return nil, errors.New(str)*/
		return nil, nil
	}
}
