package stacks

import (
	"errors"
	"fmt"
)

type frame struct {
	values map[string]any
}

func createFrame() Frame {
	out := frame{
		values: map[string]any{},
	}

	return &out
}

// Register register a value in the frame
func (app *frame) Register(name string, value any, replaceIfExists bool) error {
	if _, ok := app.values[name]; ok {
		if !replaceIfExists {
			str := fmt.Sprintf("the value (name: %s) already exists and was flagged to NOT be replaced", name)
			return errors.New(str)
		}
	}

	app.values[name] = value
	return nil
}

// Fetch fetches a value from the frame
func (app *frame) Fetch(name string) (any, error) {
	if value, ok := app.values[name]; ok {
		return value, nil
	}

	str := fmt.Sprintf("the value (name: %s) does not exists", name)
	return nil, errors.New(str)
}
