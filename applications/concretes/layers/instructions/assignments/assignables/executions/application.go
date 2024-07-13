package executions

import (
	"github.com/steve-care-software/datastencil/applications"
	"github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables"
	instruction_execution "github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables/executions"
	"github.com/steve-care-software/datastencil/domain/stacks"
)

type application struct {
	localAppBuilder   applications.LocalBuilder
	remoteAppBuilder  applications.RemoteBuilder
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
