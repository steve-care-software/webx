package services

import (
	"github.com/steve-care-software/datastencil/domain/commits/contents/actions/resources/instances"
	"github.com/steve-care-software/datastencil/domain/libraries/layers/instructions/assignments/assignables/libraries/databases/services"
	"github.com/steve-care-software/datastencil/domain/stacks"
)

type application struct {
	service           instances.Service
	assignableBuilder stacks.AssignableBuilder
}

func createApplication(
	service instances.Service,
	assignableBuilder stacks.AssignableBuilder,
) Application {
	out := application{
		service:           service,
		assignableBuilder: assignableBuilder,
	}

	return &out
}

// Execute executes the application
func (app *application) Execute(frame stacks.Frame, assignable services.Service) (stacks.Assignable, error) {
	builder := app.assignableBuilder.Create()
	if assignable.IsBegin() {
		context, err := app.service.Begin()
		if err != nil {
			return nil, err
		}

		builder.WithUnsignedInt(context)
	}

	return builder.Now()
}
