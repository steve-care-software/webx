package files

import (
	"github.com/steve-care-software/datastencil/domain/instances/layers"
)

type layerRepository struct {
	//pointerRepository pointers.Repository
	adapter layers.Adapter
}

func createLayerRepository(
	//pointerRepository pointers.Repository,
	adapter layers.Adapter,
) layers.Repository {
	out := layerRepository{
		//pointerRepository: pointerRepository,
		adapter: adapter,
	}

	return &out
}

// Retrieve retrieves a layer by path
func (app *layerRepository) Retrieve(path []string, history [][]string) (layers.Layer, error) {
	/*bytes, err := app.pointerRepository.Fetch(path, history)
	if err != nil {
		return nil, err
	}

	return app.adapter.ToInstance(bytes)*/
	return nil, nil
}
