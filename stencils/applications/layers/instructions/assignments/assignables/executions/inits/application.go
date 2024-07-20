package inits

import (
	"github.com/steve-care-software/datastencil/stencils/applications"
	instruction_inits "github.com/steve-care-software/datastencil/stencils/domain/instances/layers/instructions/assignments/assignables/executions/inits"
	"github.com/steve-care-software/datastencil/stencils/domain/stacks"
	"github.com/steve-care-software/datastencil/stencils/domain/stacks/failures"
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
func (app *application) Execute(frame stacks.Frame, executable applications.Application, assignable instruction_inits.Init) (stacks.Assignable, *uint, error) {
	pathVariable := assignable.Path()
	retPath, err := frame.FetchList(pathVariable)
	if err != nil {
		code := failures.CouldNotFetchListFromFrame
		return nil, &code, err
	}

	path := []string{}
	pathDirList := retPath.List()
	for _, oneDir := range pathDirList {
		if !oneDir.IsString() {
			code := failures.CouldNotFetchStringFromList
			return nil, &code, err
		}

		pStr := oneDir.String()
		path = append(path, *pStr)
	}

	nameVar := assignable.Name()
	name, err := frame.FetchString(nameVar)
	if err != nil {
		code := failures.CouldNotFetchStringFromFrame
		return nil, &code, err
	}

	descriptionVar := assignable.Description()
	description, err := frame.FetchString(descriptionVar)
	if err != nil {
		code := failures.CouldNotFetchStringFromFrame
		return nil, &code, err
	}

	pContext, err := executable.Init(path, name, description)
	if err != nil {
		code := failures.CouldNotExecuteInitFromExecutable
		return nil, &code, err
	}

	ins, err := app.assignableBuilder.Create().WithUnsignedInt(*pContext).Now()
	if err != nil {
		return nil, nil, err
	}

	return ins, nil, nil
}
