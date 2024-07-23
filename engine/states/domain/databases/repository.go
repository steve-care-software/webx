package databases

import (
	"github.com/steve-care-software/webx/engine/states/domain/databases/commits"
	"github.com/steve-care-software/webx/engine/states/domain/databases/pointers"
	"github.com/steve-care-software/webx/engine/states/domain/files"
)

type repository struct {
	fileRepository   files.Repository
	commitRepository commits.Repository
	pointerAdapter   pointers.Adapter
	databaseBuilder  Builder
}

func createRepository(
	fileRepository files.Repository,
	commitRepository commits.Repository,
	pointerAdapter pointers.Adapter,
	databaseBuilder Builder,
) Repository {
	out := repository{
		fileRepository:   fileRepository,
		commitRepository: commitRepository,
		pointerAdapter:   pointerAdapter,
		databaseBuilder:  databaseBuilder,
	}

	return &out
}

// Exists returns true if it exists, false otherwise
func (app *repository) Exists(path []string) bool {
	_, err := app.Retrieve(path)
	return err == nil
}

// Retrieve retrieves a database by path
func (app *repository) Retrieve(path []string) (Database, error) {
	bytes, err := app.fileRepository.RetrieveFromPath(path)
	if err != nil {
		return nil, err
	}

	pointer, err := app.pointerAdapter.ToInstance(bytes)
	if err != nil {
		return nil, err
	}

	headHash := pointer.Head()
	head, err := app.commitRepository.Retrieve(headHash)
	if err != nil {
		return nil, err
	}

	metaData := pointer.MetaData()
	return app.databaseBuilder.Create().
		WithHead(head).
		WithMetaData(metaData).
		Now()
}
