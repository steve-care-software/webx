package inputs

import (
	"bytes"
	"encoding/json"

	"github.com/steve-care-software/datastencil/stencils/domain/instances/layers/instructions/assignments/assignables/executions/executes/inputs"
)

// Adapter represents the input adapter
type Adapter struct {
	builder inputs.Builder
}

func createAdapter(
	builder inputs.Builder,
) inputs.Adapter {
	out := Adapter{
		builder: builder,
	}

	return &out
}

// ToBytes converts an instance to bytes
func (app *Adapter) ToBytes(ins inputs.Input) ([]byte, error) {
	str := app.InputToStruct(ins)
	js, err := json.Marshal(str)
	if err != nil {
		return nil, err
	}

	return js, nil
}

// ToInstance converts a bytes to instance
func (app *Adapter) ToInstance(data []byte) (inputs.Input, error) {
	ins := new(Input)
	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err := decoder.Decode(ins)
	if err != nil {
		return nil, err
	}

	return app.StructToInput(*ins)
}

// InputToStruct converts an input to struct
func (app *Adapter) InputToStruct(ins inputs.Input) Input {
	out := Input{}
	if ins.IsPath() {
		out.Path = ins.Path()
	}

	if ins.IsValue() {
		out.Value = ins.Value()
	}

	return out
}

// StructToInput converts a struct to input
func (app *Adapter) StructToInput(str Input) (inputs.Input, error) {
	builder := app.builder.Create()
	if str.Value != "" {
		builder.WithValue(str.Value)
	}

	if str.Path != "" {
		builder.WithPath(str.Path)
	}

	return builder.Now()
}
