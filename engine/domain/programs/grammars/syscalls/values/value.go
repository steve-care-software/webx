package values

import (
	"github.com/steve-care-software/webx/engine/domain/programs/grammars/blocks/lines/executions/parameters"
	"github.com/steve-care-software/webx/engine/domain/programs/grammars/blocks/lines/tokens/elements"
)

type value struct {
	parameter parameters.Parameter
	element   elements.Element
}

func createValueWithParameter(
	parameter parameters.Parameter,
) Value {
	return createValueInternally(parameter, nil)
}

func createValueWithElement(
	element elements.Element,
) Value {
	return createValueInternally(nil, element)
}

func createValueInternally(
	parameter parameters.Parameter,
	element elements.Element,
) Value {
	out := value{
		parameter: parameter,
		element:   element,
	}

	return &out
}

// IsParameter returns true if there is a parameter, false otherwise
func (obj *value) IsParameter() bool {
	return obj.parameter != nil
}

// Parameter returns the parameter, if any
func (obj *value) Parameter() parameters.Parameter {
	return obj.parameter
}

// IsElement returns true if there is an element, false otherwise
func (obj *value) IsElement() bool {
	return obj.element != nil
}

// Element returns the element, if any
func (obj *value) Element() elements.Element {
	return obj.element
}
