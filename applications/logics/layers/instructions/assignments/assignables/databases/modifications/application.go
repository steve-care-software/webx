package modifications

import (
	databases_modifications "github.com/steve-care-software/datastencil/domain/instances/databases/commits/actions/modifications"
	"github.com/steve-care-software/datastencil/domain/instances/pointers/resources/logics/layers/instructions/assignments/assignables/databases/modifications"
	"github.com/steve-care-software/datastencil/domain/stacks"
	"github.com/steve-care-software/datastencil/domain/stacks/failures"
)

type application struct {
	assignableBuilder   stacks.AssignableBuilder
	modificationBuilder databases_modifications.ModificationBuilder
}

func createApplication(
	assignableBuilder stacks.AssignableBuilder,
	modificationBuilder databases_modifications.ModificationBuilder,
) Application {
	out := application{
		assignableBuilder:   assignableBuilder,
		modificationBuilder: modificationBuilder,
	}

	return &out
}

// Execute executes the application
func (app *application) Execute(frame stacks.Frame, assignable modifications.Modification) (stacks.Assignable, *uint, error) {
	builder := app.modificationBuilder.Create()
	if assignable.IsInsert() {
		insertVar := assignable.Insert()
		insert, err := frame.FetchBytes(insertVar)
		if err != nil {
			code := failures.CouldNotFetchBytesFromFrame
			return nil, &code, nil
		}

		builder.WithInsert(insert)
	}

	if assignable.IsDelete() {
		deleteVar := assignable.Delete()
		delete, err := frame.FetchDelete(deleteVar)
		if err != nil {
			code := failures.CouldNotFetchDeleteFromFrame
			return nil, &code, nil
		}

		builder.WithDelete(delete)
	}

	modification, err := builder.Now()
	if err != nil {
		return nil, nil, err
	}

	ins, err := app.assignableBuilder.Create().WithModification(modification).Now()
	if err != nil {
		return nil, nil, err
	}

	return ins, nil, nil
}
