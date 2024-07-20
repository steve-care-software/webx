package outputs

import (
	"bytes"
	"encoding/base64"
	"encoding/json"

	"github.com/steve-care-software/datastencil/stencils/domain/instances/executions/results/success/outputs"
)

// Adapter represents an adapter
type Adapter struct {
	builder outputs.Builder
}

func createAdapter(
	builder outputs.Builder,
) outputs.Adapter {
	out := Adapter{
		builder: builder,
	}

	return &out
}

// ToBytes converts instance to bytes
func (app *Adapter) ToBytes(ins outputs.Output) ([]byte, error) {
	ptr, err := app.OutputToStruct(ins)
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
func (app *Adapter) ToInstance(data []byte) (outputs.Output, error) {
	ins := new(Output)
	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err := decoder.Decode(ins)
	if err != nil {
		return nil, err
	}

	return app.StructToOutput(*ins)
}

// OutputToStruct converts an output to struct
func (app *Adapter) OutputToStruct(ins outputs.Output) (*Output, error) {
	input := base64.StdEncoding.EncodeToString(ins.Input())
	out := Output{
		Input: input,
	}

	if ins.HasExecute() {
		execute := base64.StdEncoding.EncodeToString(ins.Execute())
		out.Execute = execute
	}

	return &out, nil
}

// StructToOutput converts a struct to output
func (app *Adapter) StructToOutput(str Output) (outputs.Output, error) {
	input, err := base64.StdEncoding.DecodeString(str.Input)
	if err != nil {
		return nil, err
	}

	builder := app.builder.Create().WithInput(input)
	if str.Execute != "" {
		execute, err := base64.StdEncoding.DecodeString(str.Execute)
		if err != nil {
			return nil, err
		}

		builder.WithExecute(execute)
	}

	return builder.Now()
}
