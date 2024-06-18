package files

import (
	"io/ioutil"
	"path/filepath"

	"github.com/steve-care-software/datastencil/domain/instances/layers"
)

type fileRepository struct {
	adapter  layers.Adapter
	basePath []string
}

func createFileRepository(
	adapter layers.Adapter,
	basePath []string,
) layers.Repository {
	out := fileRepository{
		adapter:  adapter,
		basePath: basePath,
	}

	return &out
}

// Retrieve retrieves a layer by path
func (app *fileRepository) Retrieve(path []string) (layers.Layer, error) {
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
