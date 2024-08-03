package commits

import (
	"github.com/steve-care-software/webx/engine/databases/hashes/domain/hash"
	"github.com/steve-care-software/webx/engine/states/domain/files"
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
	bytes, err := app.fileRepository.RetrieveFromPath([]string{
		hash.String(),
	})

	if err != nil {
		return nil, err
	}

	return app.adapter.ToInstance(bytes)
}
