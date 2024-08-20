package executions

import "github.com/steve-care-software/webx/engine/domain/programs/grammars/blocks/lines/executions/parameters"

type execution struct {
	fnName     string
	parameters parameters.Parameters
}

func createExecution(
	fnName string,
) Execution {
	return createExecutionInternally(fnName, nil)
}

func createExecutionWithParameters(
	fnName string,
	parameters parameters.Parameters,
) Execution {
	return createExecutionInternally(fnName, parameters)
}

func createExecutionInternally(
	fnName string,
	parameters parameters.Parameters,
) Execution {
	out := execution{
		fnName:     fnName,
		parameters: parameters,
	}

	return &out
}

// FuncName returns the func name
func (obj *execution) FuncName() string {
	return obj.fnName
}

// HasParameters returns true if there is parameters, false otherwise
func (obj *execution) HasParameters() bool {
	return obj.parameters != nil
}

// Parameters returns the parameters, if any
func (obj *execution) Parameters() parameters.Parameters {
	return obj.parameters
}
