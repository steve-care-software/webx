package jsons

import (
	"encoding/json"

	"github.com/steve-care-software/datastencil/domain/libraries"
	structs "github.com/steve-care-software/datastencil/infrastructure/jsons/structs/libraries"
)

type libraryAdapter struct {
	builder       libraries.Builder
	linksAdapter  *linksAdapter
	layersAdapter *layersAdapter
}

func createLibraryAdapter(
	builder libraries.Builder,
	linksAdapter *linksAdapter,
	layersAdapter *layersAdapter,
) libraries.Adapter {
	out := libraryAdapter{
		builder:       builder,
		linksAdapter:  linksAdapter,
		layersAdapter: layersAdapter,
	}

	return &out
}

// ToData converts a library to bytes
func (app *libraryAdapter) ToData(ins libraries.Library) ([]byte, error) {
	str := app.toStructLibrary(ins)
	data, err := json.Marshal(str)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// ToInstance converts a bytes to library
func (app *libraryAdapter) ToInstance(data []byte) (libraries.Library, error) {
	ins := structs.Library{}
	err := json.Unmarshal(data, &ins)
	if err != nil {
		return nil, err
	}

	return app.toInstanceLibrary(ins)
}

func (app *libraryAdapter) toStructLibrary(ins libraries.Library) structs.Library {
	output := structs.Library{
		Layers: app.layersAdapter.toStructLayers(ins.Layers()),
	}

	if ins.HasLinks() {
		output.Links = app.linksAdapter.toStructLinks(ins.Links())
	}

	return output
}

func (app *libraryAdapter) toInstanceLibrary(str structs.Library) (libraries.Library, error) {
	builder := app.builder.Create()
	if str.Layers != nil && len(str.Layers) > 0 {
		layers, err := app.layersAdapter.toInstanceLayers(str.Layers)
		if err != nil {
			return nil, err
		}

		builder.WithLayers(layers)

	}

	if str.Links != nil {

	}

	return builder.Now()
}
