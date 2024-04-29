package services

import (
	"github.com/steve-care-software/datastencil/applications/layers/instructions/failures"
	"github.com/steve-care-software/datastencil/domain/instances"
	"github.com/steve-care-software/datastencil/domain/instances/links/elements/layers/instructions/assignments/assignables/databases/services"
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
func (app *application) Execute(frame stacks.Frame, assignable services.Service) (stacks.Assignable, *uint, error) {
	builder := app.assignableBuilder.Create()
	if assignable.IsBegin() {
		pContext, err := app.service.Begin()
		if err != nil {
			code := failures.CouldNotBeginTransactionInDatabase
			return nil, &code, err
		}

		builder.WithUnsignedInt(*pContext)
	}

	ins, err := builder.Now()
	if err != nil {
		return nil, nil, err
	}

	return ins, nil, nil
}
