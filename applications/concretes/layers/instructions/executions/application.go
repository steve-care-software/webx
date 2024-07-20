package executions

import (
	"github.com/steve-care-software/datastencil/applications/concretes/layers/instructions/executions/merges"
	"github.com/steve-care-software/datastencil/domain/instances/layers/instructions/executions"
	"github.com/steve-care-software/datastencil/domain/stacks"
	"github.com/steve-care-software/datastencil/domain/stacks/failures"
)

type application struct {
	execMergeApp merges.Application
}

func createApplication(
	execMergeApp merges.Application,
) Application {
	out := application{
		execMergeApp: execMergeApp,
	}

	return &out
}

// Execute executes the application
func (app *application) Execute(frame stacks.Frame, assignment executions.Execution) (*uint, error) {
	executableVar := assignment.Executable()
	executable, err := frame.FetchApplication(executableVar)
	if err != nil {
		code := failures.CouldNotFetchExecutableFromFrame
		return &code, err
	}

	content := assignment.Content()
	if content.IsCommit() {
		contextVar := content.Commit()
		pContext, err := frame.FetchUnsignedInt(contextVar)
		if err != nil {
			code := failures.CouldNotFetchUnsignedIntegerFromFrame
			return &code, err
		}

		err = executable.Commit(*pContext)
		if err != nil {
			code := failures.CouldNotExecuteCommitFromExecutable
			return &code, err
		}

		return nil, nil
	}

	if content.IsRollback() {
		contextVar := content.Rollback()
		pContext, err := frame.FetchUnsignedInt(contextVar)
		if err != nil {
			code := failures.CouldNotFetchUnsignedIntegerFromFrame
			return &code, err
		}

		err = executable.Rollback(*pContext)
		if err != nil {
			code := failures.CouldNotExecuteRollbackFromExecutable
			return &code, err
		}

		return nil, nil
	}

	if content.IsCancel() {
		contextVar := content.Cancel()
		pContext, err := frame.FetchUnsignedInt(contextVar)
		if err != nil {
			code := failures.CouldNotFetchUnsignedIntegerFromFrame
			return &code, err
		}

		err = executable.Cancel(*pContext)
		if err != nil {
			code := failures.CouldNotExecuteCancelFromExecutable
			return &code, err
		}

		return nil, nil
	}

	merge := content.Merge()
	return app.execMergeApp.Execute(frame, executable, merge)
}
