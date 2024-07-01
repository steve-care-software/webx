package executions

import (
	"github.com/steve-care-software/datastencil/domain/instances/executions"
	"github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables"
	instruction_execution "github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables/executions"
	"github.com/steve-care-software/datastencil/domain/stacks"
)

type application struct {
	assignableBuilder assignables.Builder
}

func createApplication(
	assignableBuilder assignables.Builder,
) Application {
	out := application{
		assignableBuilder: assignableBuilder,
	}

	return &out
}

// Execute execute an execution instruction using the provided frame
func (app *application) Execute(frame stacks.Frame, assignable instruction_execution.Execution) (stacks.Assignable, *uint, error) {
	return nil, nil, nil
}

// ExecuteWithContext execute an execution instruction with context using the provided frame
func (app *application) ExecuteWithContext(frame stacks.Frame, assignable instruction_execution.Execution, context executions.Executions) (stacks.Assignable, *uint, error) {
	return nil, nil, nil
}
