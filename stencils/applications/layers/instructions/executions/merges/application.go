package merges

import (
	"github.com/steve-care-software/datastencil/stencils/applications"
	"github.com/steve-care-software/datastencil/stencils/domain/instances/layers/instructions/executions/merges"
	"github.com/steve-care-software/datastencil/stencils/domain/stacks"
	"github.com/steve-care-software/datastencil/stencils/domain/stacks/failures"
)

type application struct {
}

func createApplication() Application {
	out := application{}
	return &out
}

// Execute executes the application
func (app *application) Execute(frame stacks.Frame, executable applications.Application, assignment merges.Merge) (*uint, error) {
	baseContextVar := assignment.Base()
	pBaseContext, err := frame.FetchUnsignedInt(baseContextVar)
	if err != nil {
		code := failures.CouldNotFetchUnsignedIntegerFromFrame
		return &code, err
	}

	topContextVar := assignment.Top()
	pTopContext, err := frame.FetchUnsignedInt(topContextVar)
	if err != nil {
		code := failures.CouldNotFetchUnsignedIntegerFromFrame
		return &code, err
	}

	err = executable.Merge(*pBaseContext, *pTopContext)
	if err != nil {
		code := failures.CouldNotExecuteMErgeFromExecutable
		return &code, err
	}

	return nil, nil
}
