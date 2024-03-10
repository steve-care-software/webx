package jsons

import (
	"encoding/json"

	"github.com/steve-care-software/datastencil/domain/instances/libraries/layers"
	structs "github.com/steve-care-software/datastencil/infrastructure/jsons/structs/libraries/layers"
)

type layersAdapter struct {
	builder       layers.Builder
	pLayerAdapter *layerAdapter
}

func createLayersAdapter(
	builder layers.Builder,
	pLayerAdapter *layerAdapter,
) layers.Adapter {
	out := layersAdapter{
		builder:       builder,
		pLayerAdapter: pLayerAdapter,
	}

	return &out
}

// ToData convert layers to bytes
func (app *layersAdapter) ToData(ins layers.Layers) ([]byte, error) {
	str := app.toStructLayers(ins)
	data, err := json.Marshal(str)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// ToInstance convert bytes to layers
func (app *layersAdapter) ToInstance(data []byte) (layers.Layers, error) {
	ins := []structs.Layer{}
	err := json.Unmarshal(data, &ins)
	if err != nil {
		return nil, err
	}

	return app.toInstanceLayers(ins)
}

func (app *layersAdapter) toInstanceLayers(list []structs.Layer) (layers.Layers, error) {
	output := []layers.Layer{}
	for _, oneStr := range list {
		ins, err := app.pLayerAdapter.toInstanceLayer(oneStr)
		if err != nil {
			return nil, err
		}

		output = append(output, ins)
	}

	return app.builder.Create().WithList(output).Now()
}

func (app *layersAdapter) toStructLayers(ins layers.Layers) []structs.Layer {
	list := ins.List()
	output := []structs.Layer{}
	for _, oneLayer := range list {
		output = append(output, app.pLayerAdapter.toStructLayer(oneLayer))
	}

	return output
}
