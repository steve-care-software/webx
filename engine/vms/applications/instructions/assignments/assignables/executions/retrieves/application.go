package retrieves

import (
	"github.com/steve-care-software/webx/engine/stencils/applications"
	instruction_retrieves "github.com/steve-care-software/webx/engine/vms/domain/instances/layers/instructions/assignments/assignables/executions/retrieves"
	"github.com/steve-care-software/webx/engine/vms/domain/stacks"
	"github.com/steve-care-software/webx/engine/vms/domain/stacks/failures"
)

type application struct {
	assignableBuilder stacks.AssignableBuilder
}

func createApplication(
	assignableBuilder stacks.AssignableBuilder,
) Application {
	out := application{
		assignableBuilder: assignableBuilder,
	}

	return &out
}

// Execute executes the application
func (app *application) Execute(frame stacks.Frame, executable applications.Application, assignable instruction_retrieves.Retrieve) (stacks.Assignable, *uint, error) {
	contextVar := assignable.Context()
	pContext, err := frame.FetchUnsignedInt(contextVar)
	if err != nil {
		code := failures.CouldNotFetchUnsignedIntegerFromFrame
		return nil, &code, err
	}

	indexVar := assignable.Index()
	pIndex, err := frame.FetchUnsignedInt(indexVar)
	if err != nil {
		code := failures.CouldNotFetchUnsignedIntegerFromFrame
		return nil, &code, err
	}

	builder := app.assignableBuilder.Create()
	if assignable.HasLength() {
		lengthVar := assignable.Index()
		pLength, err := frame.FetchUnsignedInt(lengthVar)
		if err != nil {
			code := failures.CouldNotFetchUnsignedIntegerFromFrame
			return nil, &code, err
		}

		retExecutions, err := executable.RetrieveAll(*pContext, *pIndex, *pLength)
		if err != nil {
			code := failures.CouldNotExecuteRetrieveAllFromExecutable
			return nil, &code, err
		}

		builder.WithInstance(retExecutions)

	} else {
		retExecution, err := executable.RetrieveAt(*pContext, *pIndex)
		if err != nil {
			code := failures.CouldNotExecuteRetrieveAtFromExecutable
			return nil, &code, err
		}

		builder.WithInstance(retExecution)
	}

	ins, err := builder.Now()
	if err != nil {
		return nil, nil, err
	}

	return ins, nil, nil
}
