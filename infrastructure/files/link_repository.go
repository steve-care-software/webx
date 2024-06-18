package files

import (
	"io/ioutil"
	"path/filepath"

	"github.com/steve-care-software/datastencil/domain/instances/links"
)

type linkRepository struct {
	adapter  links.Adapter
	basePath []string
}

func createLinkRepository(
	adapter links.Adapter,
	basePath []string,
) links.Repository {
	out := linkRepository{
		adapter:  adapter,
		basePath: basePath,
	}

	return &out
}

// Retrieve retrieves a link by path
func (app *linkRepository) Retrieve(path []string) (links.Link, error) {
	fullPath := filepath.Join(append(
		app.basePath,
		path...,
	)...)

	bytes, err := ioutil.ReadFile(fullPath)
	if err != nil {
		return nil, err
	}

	return app.adapter.ToInstance(bytes)
}
