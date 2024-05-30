package success

import (
	"encoding/json"

	"github.com/steve-care-software/datastencil/domain/instances/executions/links/layers/results/success"
	json_outputs "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/commands/results/success/outputs"
	json_kinds "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/links/elements/layers/outputs/kinds"
)

// Adapter represents an adapter
type Adapter struct {
	outputAdapter *json_outputs.Adapter
	kindAdapter   *json_kinds.Adapter
	builder       success.Builder
}

func createAdapter(
	outputAdapter *json_outputs.Adapter,
	kindAdapter *json_kinds.Adapter,
	builder success.Builder,
) success.Adapter {
	out := Adapter{
		outputAdapter: outputAdapter,
		kindAdapter:   kindAdapter,
		builder:       builder,
	}

	return &out
}

// ToBytes converts instance to bytes
func (app *Adapter) ToBytes(ins success.Success) ([]byte, error) {
	ptr, err := app.SuccessToStruct(ins)
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
func (app *Adapter) ToInstance(bytes []byte) (success.Success, error) {
	ins := new(Success)
	err := json.Unmarshal(bytes, ins)
	if err != nil {
		return nil, err
	}

	return app.StructToSuccess(*ins)
}

// SuccessToStruct converts a success to struct
func (app *Adapter) SuccessToStruct(ins success.Success) (*Success, error) {
	pOutput, err := app.outputAdapter.OutputToStruct(ins.Output())
	if err != nil {
		return nil, err
	}

	pKind, err := app.kindAdapter.KindToStruct(ins.Kind())
	if err != nil {
		return nil, err
	}

	return &Success{
		Output: *pOutput,
		Kind:   *pKind,
	}, nil
}

// StructToSuccess converts a struct to success
func (app *Adapter) StructToSuccess(str Success) (success.Success, error) {
	output, err := app.outputAdapter.StructToOutput(str.Output)
	if err != nil {
		return nil, err
	}

	kind, err := app.kindAdapter.StructToKind(str.Kind)
	if err != nil {
		return nil, err
	}

	return app.builder.Create().
		WithOutput(output).
		WithKind(kind).
		Now()
}
