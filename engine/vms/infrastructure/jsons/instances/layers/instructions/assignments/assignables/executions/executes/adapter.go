package executes

import (
	"bytes"
	"encoding/json"

	"github.com/steve-care-software/webx/engine/vms/domain/instances/layers/instructions/assignments/assignables/executions/executes"
	json_inputs "github.com/steve-care-software/webx/engine/vms/infrastructure/jsons/instances/layers/instructions/assignments/assignables/executions/executes/inputs"
)

// Adapter represents an adapter
type Adapter struct {
	inputAdapter *json_inputs.Adapter
	builder      executes.Builder
}

func createAdapter(
	inputAdapter *json_inputs.Adapter,
	builder executes.Builder,
) executes.Adapter {
	out := Adapter{
		inputAdapter: inputAdapter,
		builder:      builder,
	}

	return &out
}

// ToBytes converts an instance to bytes
func (app *Adapter) ToBytes(ins executes.Execute) ([]byte, error) {
	str := app.ExecuteToStruct(ins)
	js, err := json.Marshal(str)
	if err != nil {
		return nil, err
	}

	return js, nil
}

// ToInstance converts a bytes to instance
func (app *Adapter) ToInstance(data []byte) (executes.Execute, error) {
	ins := new(Execute)
	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err := decoder.Decode(ins)
	if err != nil {
		return nil, err
	}

	return app.StructToExecute(*ins)
}

// ExecuteToStruct converts an execute to struct
func (app *Adapter) ExecuteToStruct(ins executes.Execute) Execute {
	input := app.inputAdapter.InputToStruct(ins.Input())
	out := Execute{
		Context: ins.Context(),
		Input:   input,
	}

	if ins.HasLayer() {
		out.Layer = ins.Layer()
	}

	return out
}

// StructToInput converts a struct to input
func (app *Adapter) StructToExecute(str Execute) (executes.Execute, error) {
	input, err := app.inputAdapter.StructToInput(str.Input)
	if err != nil {
		return nil, err
	}

	builder := app.builder.Create().
		WithContext(str.Context).
		WithInput(input)

	if str.Layer != "" {
		builder.WithLayer(str.Layer)
	}

	return builder.Now()
}
