package executes

import (
	"github.com/steve-care-software/datastencil/stencils/applications"
	instruction_executes "github.com/steve-care-software/datastencil/stencils/domain/instances/layers/instructions/assignments/assignables/executions/executes"
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
func (app *application) Execute(frame stacks.Frame, executable applications.Application, assignable instruction_executes.Execute) (stacks.Assignable, *uint, error) {
	contextVar := assignable.Context()
	pContext, err := frame.FetchUnsignedInt(contextVar)
	if err != nil {
		code := failures.CouldNotFetchUnsignedIntegerFromFrame
		return nil, &code, err
	}

	layerPath := []string{}
	if assignable.HasLayer() {
		layerPathVariable := assignable.Layer()
		retLayerPath, err := frame.FetchList(layerPathVariable)
		if err != nil {
			code := failures.CouldNotFetchListFromFrame
			return nil, &code, err
		}

		pathDirList := retLayerPath.List()
		for _, oneDir := range pathDirList {
			if !oneDir.IsString() {
				code := failures.CouldNotFetchStringFromList
				return nil, &code, err
			}

			pStr := oneDir.String()
			layerPath = append(layerPath, *pStr)
		}
	}

	input := assignable.Input()
	assignableBuilder := app.assignableBuilder.Create()
	if input.IsValue() {
		valueVar := input.Value()
		inputBytes, err := frame.FetchBytes(valueVar)
		if err != nil {
			code := failures.CouldNotFetchBytesFromFrame
			return nil, &code, err
		}

		if len(layerPath) <= 0 {
			retBytes, err := executable.Execute(*pContext, inputBytes)
			if err != nil {
				code := failures.CouldNotExecuteExecuteFromExecutable
				return nil, &code, err
			}

			assignableBuilder.WithBytes(retBytes)
		} else {
			retBytes, err := executable.ExecuteLayer(*pContext, inputBytes, layerPath)
			if err != nil {
				code := failures.CouldNotExecuteExecuteLayerFromExecutable
				return nil, &code, err
			}

			assignableBuilder.WithBytes(retBytes)
		}
	}

	if input.IsPath() {
		pathVariable := input.Path()
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

		if len(layerPath) <= 0 {
			retBytes, err := executable.ExecuteWithPath(*pContext, path)
			if err != nil {
				code := failures.CouldNotExecuteWithPathFromExecutable
				return nil, &code, err
			}

			assignableBuilder.WithBytes(retBytes)
		} else {
			retBytes, err := executable.ExecuteLayerWithPath(*pContext, path, layerPath)
			if err != nil {
				code := failures.CouldNotExecuteExecuteLayerWithPathFromExecutable
				return nil, &code, err
			}

			assignableBuilder.WithBytes(retBytes)
		}
	}

	ins, err := assignableBuilder.Now()
	if err != nil {
		return nil, nil, err
	}

	return ins, nil, nil
}
