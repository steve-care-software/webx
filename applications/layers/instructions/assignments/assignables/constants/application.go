package constants

import (
	"github.com/steve-care-software/datastencil/domain/libraries/layers/instructions/assignments/assignables/constants"
	"github.com/steve-care-software/datastencil/domain/stacks"
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
func (app *application) Execute(frame stacks.Frame, assignable constants.Constant) (stacks.Assignable, error) {
	builder := app.assignableBuilder.Create()
	if assignable.IsBool() {
		pBool := assignable.Bool()
		builder.WithBool(*pBool)
	}

	if assignable.IsBytes() {
		bytes := assignable.Bytes()
		builder.WithBytes(bytes)
	}

	return builder.Now()
}
