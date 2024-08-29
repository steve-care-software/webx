package parameters

import (
	"github.com/steve-care-software/webx/engine/domain/programs/grammars/blocks/lines/executions/parameters/values"
)

type parameter struct {
	name  string
	value values.Value
}

func createParameter(
	name string,
	value values.Value,
) Parameter {
	out := parameter{
		name:  name,
		value: value,
	}

	return &out
}

// Name returns the name
func (obj *parameter) Name() string {
	return obj.name
}

// Value returns the value
func (obj *parameter) Value() values.Value {
	return obj.value
}
