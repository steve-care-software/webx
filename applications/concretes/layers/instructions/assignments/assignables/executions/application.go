package executions

import (
	"github.com/steve-care-software/datastencil/applications/concretes/layers/instructions/assignments/assignables/executions/amounts"
	"github.com/steve-care-software/datastencil/applications/concretes/layers/instructions/assignments/assignables/executions/begins"
	"github.com/steve-care-software/datastencil/applications/concretes/layers/instructions/assignments/assignables/executions/executes"
	"github.com/steve-care-software/datastencil/applications/concretes/layers/instructions/assignments/assignables/executions/heads"
	"github.com/steve-care-software/datastencil/applications/concretes/layers/instructions/assignments/assignables/executions/inits"
	"github.com/steve-care-software/datastencil/applications/concretes/layers/instructions/assignments/assignables/executions/retrieves"
	instruction_execution "github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables/executions"
	"github.com/steve-care-software/datastencil/domain/stacks"
	"github.com/steve-care-software/datastencil/domain/stacks/failures"
)

type application struct {
	execAmountApp      amounts.Application
	execBeginApp       begins.Application
	execExecuteApp     executes.Application
	execHeadApp        heads.Application
	execInitApp        inits.Application
	execRetrieveApp    retrieves.Application
	assignableBuilder  stacks.AssignableBuilder
	assignablesBuilder stacks.AssignablesBuilder
}

func createApplication(
	execAmountApp amounts.Application,
	execBeginApp begins.Application,
	execExecuteApp executes.Application,
	execHeadApp heads.Application,
	execInitApp inits.Application,
	execRetrieveApp retrieves.Application,
	assignableBuilder stacks.AssignableBuilder,
	assignablesBuilder stacks.AssignablesBuilder,
) Application {
	out := application{
		execAmountApp:      execAmountApp,
		execBeginApp:       execBeginApp,
		execExecuteApp:     execExecuteApp,
		execHeadApp:        execHeadApp,
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
		code := failures.CouldNotFetchExecutableeFromFrame
		return nil, &code, err
	}

	content := assignable.Content()
	if content.IsInit() {
		init := content.Init()
		return app.execInitApp.Execute(frame, retExecutable, init)
	}

	if content.IsBegin() {
		begin := content.Begin()
		return app.execBeginApp.Execute(frame, retExecutable, begin)
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
		amount := content.Amount()
		return app.execAmountApp.Execute(frame, retExecutable, amount)
	}

	if content.IsHead() {
		head := content.Head()
		return app.execHeadApp.Execute(frame, retExecutable, head)
	}

	retPathList, err := retExecutable.List()
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

	return ins, nil, nil
}
