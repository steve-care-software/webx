package commits

import (
	"github.com/steve-care-software/datastencil/states/domain/files"
	"github.com/steve-care-software/datastencil/states/domain/hash"
)

type repository struct {
	adapter        Adapter
	fileRepository files.Repository
}

func createRepository(
	adapter Adapter,
	fileRepository files.Repository,
) Repository {
	out := repository{
		adapter:        adapter,
		fileRepository: fileRepository,
	}

	return &out
}

// Retrieve retrieves a commit by hash
func (app *repository) Retrieve(hash hash.Hash) (Commit, error) {
	bytes, err := app.fileRepository.Retrieve([]string{
		hash.String(),
	})

	if err != nil {
		return nil, err
	}

	return app.adapter.ToInstance(bytes)
}
