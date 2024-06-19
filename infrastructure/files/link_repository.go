package files

import (
	"github.com/steve-care-software/datastencil/domain/instances/links"
	"github.com/steve-care-software/datastencil/domain/instances/pointers"
)

type linkRepository struct {
	pointerRepository pointers.Repository
	adapter           links.Adapter
}

func createLinkRepository(
	pointerRepository pointers.Repository,
	adapter links.Adapter,
) links.Repository {
	out := linkRepository{
		pointerRepository: pointerRepository,
		adapter:           adapter,
	}

	return &out
}

// Retrieve retrieves a link by path
func (app *linkRepository) Retrieve(path []string, history [][]string) (links.Link, error) {
	bytes, err := app.pointerRepository.Fetch(path, history)
	if err != nil {
		return nil, err
	}

	return app.adapter.ToInstance(bytes)
}
