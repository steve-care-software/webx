package assignments

import (
	"github.com/steve-care-software/datastencil/applications/layers/instructions/assignments/assignables"
	"github.com/steve-care-software/datastencil/domain/libraries/layers/instructions/assignments"
	"github.com/steve-care-software/datastencil/domain/stacks"
)

type application struct {
	execAssignableApp assignables.Application
	assignmentBuilder stacks.AssignmentBuilder
}

func createApplication(
	execAssignableApp assignables.Application,
	assignmentBuilder stacks.AssignmentBuilder,
) Application {
	out := application{
		execAssignableApp: execAssignableApp,
		assignmentBuilder: assignmentBuilder,
	}

	return &out
}

// Execute executes the application
func (app *application) Execute(frame stacks.Frame, assignment assignments.Assignment) (stacks.Assignment, error) {
	assignable := assignment.Assignable()
	retAssignable, err := app.execAssignableApp.Execute(frame, assignable)
	if err != nil {
		return nil, err
	}

	name := assignment.Name()
	return app.assignmentBuilder.Create().
		WithName(name).
		WithAssignable(retAssignable).
		Now()
}
