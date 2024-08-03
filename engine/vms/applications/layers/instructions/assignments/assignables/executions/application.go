package executions

import (
	"github.com/steve-care-software/webx/engine/vms/applications/layers/instructions/assignments/assignables/executions/executes"
	"github.com/steve-care-software/webx/engine/vms/applications/layers/instructions/assignments/assignables/executions/inits"
	"github.com/steve-care-software/webx/engine/vms/applications/layers/instructions/assignments/assignables/executions/retrieves"
	instruction_execution "github.com/steve-care-software/webx/engine/vms/domain/instances/layers/instructions/assignments/assignables/executions"
	"github.com/steve-care-software/webx/engine/vms/domain/stacks"
	"github.com/steve-care-software/webx/engine/vms/domain/stacks/failures"
)

type application struct {
	execExecuteApp     executes.Application
	execInitApp        inits.Application
	execRetrieveApp    retrieves.Application
	assignableBuilder  stacks.AssignableBuilder
	assignablesBuilder stacks.AssignablesBuilder
}

func createApplication(
	execExecuteApp executes.Application,
	execInitApp inits.Application,
	execRetrieveApp retrieves.Application,
	assignableBuilder stacks.AssignableBuilder,
	assignablesBuilder stacks.AssignablesBuilder,
) Application {
	out := application{
		execExecuteApp:     execExecuteApp,
		execInitApp:        execInitApp,
		execRetrieveApp:    execRetrieveApp,
		assignableBuilder:  assignableBuilder,
		assignablesBuilder: assignablesBuilder,
	}

	return &out
}

// Execute execute an execution instruction using the provided frame
func (app *application) Execute(frame stacks.Frame, assignable instruction_execution.Execution) (stacks.Assignable, *uint, error) {
	executableVar := assignable.Executable()
	retExecutable, err := frame.FetchApplication(executableVar)
	if err != nil {
		code := failures.CouldNotFetchExecutableFromFrame
		return nil, &code, err
	}

	content := assignable.Content()
	if content.IsInit() {
		init := content.Init()
		return app.execInitApp.Execute(frame, retExecutable, init)
	}

	if content.IsBegin() {
		/*pathVar := content.Begin()
		retPath, err := frame.FetchList(pathVar)
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

		pContext, err := retExecutable.Begin(path)
		if err != nil {
			code := failures.CouldNotExecuteBeginFromExecutable
			return nil, &code, err
		}

		ins, err := app.assignableBuilder.Create().WithUnsignedInt(*pContext).Now()
		if err != nil {
			return nil, nil, err
		}

		return ins, nil, nil*/
		return nil, nil, nil
	}

	if content.IsExecute() {
		execute := content.Execute()
		return app.execExecuteApp.Execute(frame, retExecutable, execute)
	}

	if content.IsRetrieve() {
		retrieve := content.Retrieve()
		return app.execRetrieveApp.Execute(frame, retExecutable, retrieve)
	}

	if content.IsAmount() {
		/*amount := content.Amount()
		pContext, err := frame.FetchUnsignedInt(amount)
		if err != nil {
			code := failures.CouldNotFetchUnsignedIntegerFromFrame
			return nil, &code, err
		}

		pAmount, err := retExecutable.Amount(*pContext)
		if err != nil {
			code := failures.CouldNotExecuteAmountFromExecutable
			return nil, &code, err
		}

		ins, err := app.assignableBuilder.Create().WithUnsignedInt(*pAmount).Now()
		if err != nil {
			return nil, nil, err
		}

		return ins, nil, nil*/
		return nil, nil, nil
	}

	if content.IsHead() {
		/*amount := content.Head()
		pContext, err := frame.FetchUnsignedInt(amount)
		if err != nil {
			code := failures.CouldNotFetchUnsignedIntegerFromFrame
			return nil, &code, err
		}

		pAmount, err := retExecutable.Amount(*pContext)
		if err != nil {
			code := failures.CouldNotExecuteHeadFromExecutable
			return nil, &code, err
		}

		ins, err := app.assignableBuilder.Create().WithUnsignedInt(*pAmount).Now()
		if err != nil {
			return nil, nil, err
		}

		return ins, nil, nil*/
		return nil, nil, nil
	}

	/*retPathList, err := retExecutable.List()
	if err != nil {
		code := failures.CouldNotExecuteListFromExecutable
		return nil, &code, err
	}

	list := []stacks.Assignable{}
	for _, onePath := range retPathList {
		dirList := []stacks.Assignable{}
		for _, oneDir := range onePath {
			ins, err := app.assignableBuilder.Create().WithString(oneDir).Now()
			if err != nil {
				return nil, nil, err
			}

			dirList = append(dirList, ins)
		}

		assignables, err := app.assignablesBuilder.Create().WithList(dirList).Now()
		if err != nil {
			return nil, nil, err
		}

		ins, err := app.assignableBuilder.Create().WithList(assignables).Now()
		if err != nil {
			return nil, nil, err
		}

		list = append(list, ins)
	}

	assignables, err := app.assignablesBuilder.Create().WithList(list).Now()
	if err != nil {
		return nil, nil, err
	}

	ins, err := app.assignableBuilder.Create().WithList(assignables).Now()
	if err != nil {
		return nil, nil, err
	}

	return ins, nil, nil*/
	return nil, nil, nil
}
