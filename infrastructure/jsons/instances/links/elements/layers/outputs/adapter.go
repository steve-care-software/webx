package outputs

import (
	"encoding/json"

	"github.com/steve-care-software/datastencil/domain/instances/links/elements/layers/outputs"
	json_kinds "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/links/elements/layers/outputs/kinds"
)

// Adapter represents an adapter
type Adapter struct {
	kindAdapter *json_kinds.Adapter
	builder     outputs.Builder
}

func createAdapter(
	kindAdapter *json_kinds.Adapter,
	builder outputs.Builder,
) outputs.Adapter {
	out := Adapter{
		kindAdapter: kindAdapter,
		builder:     builder,
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
func (app *Adapter) ToInstance(bytes []byte) (outputs.Output, error) {
	ins := new(Output)
	err := json.Unmarshal(bytes, ins)
	if err != nil {
		return nil, err
	}

	return app.StructToOutput(*ins)
}

// OutputToStruct converts an output to struct
func (app *Adapter) OutputToStruct(ins outputs.Output) (*Output, error) {
	ptr, err := app.kindAdapter.KindToStruct(ins.Kind())
	if err != nil {
		return nil, err
	}

	out := Output{
		Variable: ins.Variable(),
		Kind:     *ptr,
	}

	if ins.HasExecute() {
		out.Execute = ins.Execute()
	}

	return &out, nil
}

// StructToOutput converts a struct to output
func (app *Adapter) StructToOutput(str Output) (outputs.Output, error) {
	ins, err := app.kindAdapter.StructToKind(str.Kind)
	if err != nil {
		return nil, err
	}

	builder := app.builder.Create().WithVariable(str.Variable).WithKind(ins)
	if str.Execute != nil {
		builder.WithExecute(str.Execute)
	}

	return builder.Now()
}
