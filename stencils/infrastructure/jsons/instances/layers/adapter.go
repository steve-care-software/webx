package layers

import (
	"bytes"
	"encoding/json"

	"github.com/steve-care-software/datastencil/stencils/domain/instances/layers"
	json_instructions "github.com/steve-care-software/datastencil/stencils/infrastructure/jsons/instances/layers/instructions"
	json_output "github.com/steve-care-software/datastencil/stencils/infrastructure/jsons/instances/layers/outputs"
	json_references "github.com/steve-care-software/datastencil/stencils/infrastructure/jsons/instances/layers/references"
)

// Adapter represents the adapter
type Adapter struct {
	instructionsAdapter *json_instructions.Adapter
	outputAdapter       *json_output.Adapter
	referenceAdapter    *json_references.Adapter
	builder             layers.Builder
}

func createAdapter(
	instructionsAdapter *json_instructions.Adapter,
	outputAdapter *json_output.Adapter,
	referenceAdapter *json_references.Adapter,
	builder layers.Builder,
) layers.Adapter {
	out := Adapter{
		instructionsAdapter: instructionsAdapter,
		outputAdapter:       outputAdapter,
		referenceAdapter:    referenceAdapter,
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
func (app *Adapter) ToInstance(data []byte) (layers.Layer, error) {
	ins := new(Layer)
	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err := decoder.Decode(ins)
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

	output := Layer{
		Instructions: *&ptrInstructions,
		Output:       *ptrOutput,
		Input:        ins.Input(),
	}

	if ins.HasReferences() {
		referencesList, err := app.referenceAdapter.ReferencesToStruct(ins.References())
		if err != nil {
			return nil, err
		}

		output.References = referencesList
	}

	return &output, nil
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

	builder := app.builder.Create().
		WithInstructions(instructions).
		WithOutput(output).
		WithInput(str.Input)

	if str.References != nil && len(str.References) > 0 {
		references, err := app.referenceAdapter.StructToReferences(str.References)
		if err != nil {
			return nil, err
		}

		builder.WithReferences(references)
	}

	return builder.Now()
}
