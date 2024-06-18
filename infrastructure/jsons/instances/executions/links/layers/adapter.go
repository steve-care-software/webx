package layers

import (
	"encoding/base64"
	"encoding/json"

	"github.com/steve-care-software/datastencil/domain/instances/executions/links/layers"
	json_results "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/executions/links/layers/results"
	json_layers "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/layers"
)

// Adapter represents a layer adapter
type Adapter struct {
	layerAdapter  *json_layers.Adapter
	resultAdapter *json_results.Adapter
	builder       layers.Builder
	layerBuilder  layers.LayerBuilder
}

func createAdapter(
	layerAdapter *json_layers.Adapter,
	resultAdapter *json_results.Adapter,
	builder layers.Builder,
	layerBuilder layers.LayerBuilder,
) layers.Adapter {
	out := Adapter{
		layerAdapter:  layerAdapter,
		resultAdapter: resultAdapter,
		builder:       builder,
		layerBuilder:  layerBuilder,
	}

	return &out
}

// ToBytes converts instance to bytes
func (app *Adapter) ToBytes(ins layers.Layers) ([]byte, error) {
	ptr, err := app.LayersToStruct(ins)
	if err != nil {
		return nil, err
	}

	js, err := json.Marshal(ptr)
	if err != nil {
		return nil, err
	}

	return js, nil
}

// ToInstance converts bytes to instance
func (app *Adapter) ToInstance(bytes []byte) (layers.Layers, error) {
	ins := new([]Layer)
	err := json.Unmarshal(bytes, ins)
	if err != nil {
		return nil, err
	}

	return app.StructToLayers(*ins)
}

// LayersToStruct converts a layers to struct
func (app *Adapter) LayersToStruct(ins layers.Layers) ([]Layer, error) {
	out := []Layer{}
	list := ins.List()
	for _, oneIns := range list {
		ptr, err := app.LayerToStruct(oneIns)
		if err != nil {
			return nil, err
		}

		out = append(out, *ptr)
	}

	return out, nil
}

// StructToLayers converts a struct to layers
func (app *Adapter) StructToLayers(str []Layer) (layers.Layers, error) {
	list := []layers.Layer{}
	for _, oneStruct := range str {
		ins, err := app.StructToLayer(oneStruct)
		if err != nil {
			return nil, err
		}

		list = append(list, ins)
	}

	return app.builder.Create().
		WithList(list).
		Now()
}

// LayerToStruct converts a layer to struct
func (app *Adapter) LayerToStruct(ins layers.Layer) (*Layer, error) {
	ptrSource, err := app.layerAdapter.LayerToStruct(ins.Source())
	if err != nil {
		return nil, err
	}

	ptrResult, err := app.resultAdapter.ResultToStruct(ins.Result())
	if err != nil {
		return nil, err
	}

	encoded := base64.StdEncoding.EncodeToString(ins.Input())
	return &Layer{
		Source: *ptrSource,
		Result: *ptrResult,
		Input:  encoded,
	}, nil
}

// StructToLayer converts a struct to layer
func (app *Adapter) StructToLayer(str Layer) (layers.Layer, error) {
	sourceIns, err := app.layerAdapter.StructToLayer(str.Source)
	if err != nil {
		return nil, err
	}

	resultIns, err := app.resultAdapter.StructToResult(str.Result)
	if err != nil {
		return nil, err
	}

	decoded, err := base64.StdEncoding.DecodeString(str.Input)
	if err != nil {
		return nil, err
	}

	return app.layerBuilder.Create().
		WithInput(decoded).
		WithSource(sourceIns).
		WithResult(resultIns).
		Now()
}
