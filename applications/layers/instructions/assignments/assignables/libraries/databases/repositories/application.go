package repositories

import (
	"github.com/steve-care-software/datastencil/domain/commits/contents/actions/resources/instances"
	"github.com/steve-care-software/datastencil/domain/commits/contents/actions/resources/instances/skeletons"
	"github.com/steve-care-software/datastencil/domain/libraries/layers/instructions/assignments/assignables/libraries/databases/repositories"
	"github.com/steve-care-software/datastencil/domain/stacks"
	stacks_databases "github.com/steve-care-software/datastencil/domain/stacks/databases"
	stacks_libraries "github.com/steve-care-software/datastencil/domain/stacks/libraries"
)

type application struct {
	repository        instances.Repository
	libraryBuilder    stacks_libraries.Builder
	databaseBuilder   stacks_databases.Builder
	assignableBuilder stacks.AssignableBuilder
	skeleton          skeletons.Skeleton
}

func createApplication(
	repository instances.Repository,
	libraryBuilder stacks_libraries.Builder,
	databaseBuilder stacks_databases.Builder,
	assignableBuilder stacks.AssignableBuilder,
	skeleton skeletons.Skeleton,
) Application {
	out := application{
		repository:        repository,
		libraryBuilder:    libraryBuilder,
		databaseBuilder:   databaseBuilder,
		assignableBuilder: assignableBuilder,
		skeleton:          skeleton,
	}

	return &out
}

// Execute executes the application
func (app *application) Execute(frame stacks.Frame, assignable repositories.Repository) (stacks.Assignable, error) {
	builder := app.assignableBuilder.Create()
	if assignable.IsSkeleton() {
		database, err := app.databaseBuilder.Create().
			WithSkeleton(app.skeleton).
			Now()

		if err != nil {
			return nil, err
		}

		builder.WithDatabase(database)
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

		library, err := app.libraryBuilder.Create().
			WithInstance(retIns).
			Now()

		if err != nil {
			return nil, err
		}

		builder.WithLibrary(library)
	}

	return builder.Now()
}
