package repositories

import (
	"github.com/steve-care-software/datastencil/applications/layers/instructions/failures"
	"github.com/steve-care-software/datastencil/domain/instances"
	"github.com/steve-care-software/datastencil/domain/instances/links/layers/instructions/assignments/assignables/databases/repositories"
	"github.com/steve-care-software/datastencil/domain/instances/skeletons"
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
func (app *application) Execute(frame stacks.Frame, assignable repositories.Repository) (stacks.Assignable, *uint, error) {
	builder := app.assignableBuilder.Create()
	if assignable.IsSkeleton() {
		builder.WithSkeleton(app.skeleton)
	}

	if assignable.IsHeight() {
		pHeight, err := app.repository.Height()
		if err != nil {
			return nil, nil, err
		}

		builder.WithUnsignedInt(*pHeight)
	}

	if assignable.IsList() {
		listVar := assignable.List()
		query, err := frame.FetchQuery(listVar)
		if err != nil {
			code := failures.CouldNotFetchListQueryFromFrame
			return nil, &code, err
		}

		retHashList, err := app.repository.List(query)
		if err != nil {
			code := failures.CouldNotListInstancesFromDatabase
			return nil, &code, err
		}

		builder.WithHashList(retHashList)

	}

	if assignable.IsRetrieve() {
		retrieveVar := assignable.Retrieve()
		query, err := frame.FetchQuery(retrieveVar)
		if err != nil {
			code := failures.CouldNotFetchRetrieveQueryFromFrame
			return nil, &code, err
		}

		retIns, err := app.repository.Retrieve(query)
		if err != nil {
			code := failures.CouldNotRetrieveInstanceFromDatabase
			return nil, &code, err
		}

		builder.WithInstance(retIns)
	}

	ins, err := builder.Now()
	if err != nil {
		return nil, nil, err
	}

	return ins, nil, nil
}
