package layers

import (
	"encoding/json"

	"github.com/steve-care-software/datastencil/domain/instances/links/elements/layers"
	json_instructions "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/links/elements/layers/instructions"
	json_output "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/links/elements/layers/outputs"
)

// Adapter represents the adapter
type Adapter struct {
	instructionsAdapter *json_instructions.Adapter
	outputAdapter       *json_output.Adapter
	layerBuilder        layers.LayerBuilder
	builder             layers.Builder
}

func createAdapter(
	instructionsAdapter *json_instructions.Adapter,
	outputAdapter *json_output.Adapter,
	layerBuilder layers.LayerBuilder,
	builder layers.Builder,
) layers.Adapter {
	out := Adapter{
		instructionsAdapter: instructionsAdapter,
		outputAdapter:       outputAdapter,
		layerBuilder:        layerBuilder,
		builder:             builder,
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
	output := []Layer{}
	list := ins.List()
	for _, oneIns := range list {
		ptr, err := app.LayerToStruct(oneIns)
		if err != nil {
			return nil, err
		}

		output = append(output, *ptr)
	}

	return output, nil
}

// StructToLayers converts a struct to layers
func (app *Adapter) StructToLayers(list []Layer) (layers.Layers, error) {
	output := []layers.Layer{}
	for _, oneStr := range list {
		ins, err := app.StructToLayer(oneStr)
		if err != nil {
			return nil, err
		}

		output = append(output, ins)
	}

	return app.builder.Create().
		WithList(output).
		Now()
}

// LayerToStruct converts a layer to struct
func (app *Adapter) LayerToStruct(ins layers.Layer) (*Layer, error) {
	ptrInstructions, err := app.instructionsAdapter.InstructionsToStruct(ins.Instructions())
	if err != nil {
		return nil, err
	}

	ptrOutput, err := app.outputAdapter.OutputToStruct(ins.Output())
	if err != nil {
		return nil, err
	}

	return &Layer{
		Instructions: *&ptrInstructions,
		Output:       *ptrOutput,
		Input:        ins.Input(),
	}, nil
}

// StructToLayer converts a struct to layer
func (app *Adapter) StructToLayer(str Layer) (layers.Layer, error) {
	instructions, err := app.instructionsAdapter.StructToInstructions(str.Instructions)
	if err != nil {
		return nil, err
	}

	output, err := app.outputAdapter.StructToOutput(str.Output)
	if err != nil {
		return nil, err
	}

	return app.layerBuilder.Create().
		WithInstructions(instructions).
		WithOutput(output).
		WithInput(str.Input).
		Now()
}
