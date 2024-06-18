package layers

import (
	"encoding/json"

	"github.com/steve-care-software/datastencil/domain/instances/layers"
	json_instructions "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/layers/instructions"
	json_output "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/layers/outputs"
)

// Adapter represents the adapter
type Adapter struct {
	instructionsAdapter *json_instructions.Adapter
	outputAdapter       *json_output.Adapter
	builder             layers.Builder
}

func createAdapter(
	instructionsAdapter *json_instructions.Adapter,
	outputAdapter *json_output.Adapter,
	builder layers.Builder,
) layers.Adapter {
	out := Adapter{
		instructionsAdapter: instructionsAdapter,
		outputAdapter:       outputAdapter,
		builder:             builder,
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

	return app.builder.Create().
		WithInstructions(instructions).
		WithOutput(output).
		WithInput(str.Input).
		Now()
}
