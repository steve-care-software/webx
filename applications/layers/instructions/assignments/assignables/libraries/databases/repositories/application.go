package repositories

import (
	"github.com/steve-care-software/datastencil/domain/instances"
	"github.com/steve-care-software/datastencil/domain/instances/libraries/layers/instructions/assignments/assignables/libraries/databases/repositories"
	"github.com/steve-care-software/datastencil/domain/skeletons"
	"github.com/steve-care-software/datastencil/domain/stacks"
)

type application struct {
	repository        instances.Repository
	assignableBuilder stacks.AssignableBuilder
	skeleton          skeletons.Skeleton
}

func createApplication(
	repository instances.Repository,
	assignableBuilder stacks.AssignableBuilder,
	skeleton skeletons.Skeleton,
) Application {
	out := application{
		repository:        repository,
		assignableBuilder: assignableBuilder,
		skeleton:          skeleton,
	}

	return &out
}

// Execute executes the application
func (app *application) Execute(frame stacks.Frame, assignable repositories.Repository) (stacks.Assignable, error) {
	builder := app.assignableBuilder.Create()
	if assignable.IsSkeleton() {
		builder.WithSkeleton(app.skeleton)
	}

	if assignable.IsHeight() {
		height, err := app.repository.Height()
		if err != nil {
			return nil, err
		}

		builder.WithUnsignedInt(height)
	}

	if assignable.IsList() {
		listVar := assignable.List()
		query, err := frame.FetchQuery(listVar)
		if err != nil {
			return nil, err
		}

		retHashList, err := app.repository.List(query)
		if err != nil {
			return nil, err
		}

		builder.WithHashList(retHashList)

	}

	if assignable.IsRetrieve() {
		retrieveVar := assignable.Retrieve()
		query, err := frame.FetchQuery(retrieveVar)
		if err != nil {
			return nil, err
		}

		retIns, err := app.repository.Retrieve(query)
		if err != nil {
			return nil, err
		}

		builder.WithInstance(retIns)
	}

	return builder.Now()
}
