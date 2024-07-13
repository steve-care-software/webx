package layers

import (
	"encoding/base64"
	"encoding/json"

	"github.com/steve-care-software/datastencil/domain/instances/executions/layers"
	json_results "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/executions/layers/results"
	json_layers "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/layers"
)

// Adapter represents a layer adapter
type Adapter struct {
	layerAdapter  *json_layers.Adapter
	resultAdapter *json_results.Adapter
	builder       layers.Builder
}

func createAdapter(
	layerAdapter *json_layers.Adapter,
	resultAdapter *json_results.Adapter,
	builder layers.Builder,
) layers.Adapter {
	out := Adapter{
		layerAdapter:  layerAdapter,
		resultAdapter: resultAdapter,
		builder:       builder,
	}

	return &out
}

// ToBytes converts instance to bytes
func (app *Adapter) ToBytes(ins layers.Layer) ([]byte, error) {
	ptr, err := app.LayerToStruct(ins)
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
func (app *Adapter) ToInstance(bytes []byte) (layers.Layer, error) {
	ins := new(Layer)
	err := json.Unmarshal(bytes, ins)
	if err != nil {
		return nil, err
	}

	return app.StructToLayer(*ins)
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

	return app.builder.Create().
		WithInput(decoded).
		WithSource(sourceIns).
		WithResult(resultIns).
		Now()
}
