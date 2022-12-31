package vms

import (
	"errors"
	"fmt"
	"strings"

	"github.com/steve-care-software/webx/programs/domain/programs/modules"
)

type moduleContainers struct {
}

func createModuleContainers() *moduleContainers {
	out := moduleContainers{}
	return &out
}

// Execute executes the application
func (app *moduleContainers) Execute() map[uint]modules.ExecuteFn {
	mapWithKeynames := app.containerMapWithStringKeynames()
	list := app.containerList()
	return map[uint]modules.ExecuteFn{
		ModuleContainerMapWithKeynames: mapWithKeynames,
		ModuleContainerList:            list,
	}
}

func (app *moduleContainers) containerMapWithStringKeynames() modules.ExecuteFn {
	return func(input map[uint]interface{}) (interface{}, error) {
		output := map[string]interface{}{}
		if name, ok := input[0].(string); ok {
			if value, ok := input[1]; ok {
				name = strings.TrimSpace(name)
				output[name] = value
				return output, nil
			}

			str := fmt.Sprintf("the value was expected to be declared")
			return nil, errors.New(str)
		}

		str := fmt.Sprintf("the name was expected to be declared and contain a string")
		return nil, errors.New(str)
	}
}

func (app *moduleContainers) containerList() modules.ExecuteFn {
	return func(input map[uint]interface{}) (interface{}, error) {
		values := []interface{}{}
		for keyname, element := range input {
			indexKeyname := uint(len(values))
			if keyname != indexKeyname {
				continue
			}

			values = append(values, element)
		}

		return values, nil
	}
}
