package executions

import "github.com/steve-care-software/webx/engine/domain/grammars/tokens/elements"

type execution struct {
	fnName   string
	elements elements.Elements
}

func createExecution(
	fnName string,
) Execution {
	return createExecutionInternally(fnName, nil)
}

func createExecutionWithElements(
	fnName string,
	elements elements.Elements,
) Execution {
	return createExecutionInternally(fnName, elements)
}

func createExecutionInternally(
	fnName string,
	elements elements.Elements,
) Execution {
	out := execution{
		fnName:   fnName,
		elements: elements,
	}

	return &out
}

// FuncName returns the func name
func (obj *execution) FuncName() string {
	return obj.fnName
}

// HasElements returns true if there is elements, false otherwise
func (obj *execution) HasElements() bool {
	return obj.elements != nil
}

// Elements returns the elements
func (obj *execution) Elements() elements.Elements {
	return obj.elements
}
