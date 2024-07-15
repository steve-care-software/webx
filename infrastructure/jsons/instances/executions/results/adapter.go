package results

import (
	"encoding/json"

	"github.com/steve-care-software/datastencil/domain/instances/executions/results"
	json_interruptions "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/executions/results/interruptions"
	json_success "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/executions/results/success"
)

// Adapter represents an adapter
type Adapter struct {
	interruptionAdapter *json_interruptions.Adapter
	successAdapter      *json_success.Adapter
	builder             results.Builder
}

func createAdapter(
	interruptionAdapter *json_interruptions.Adapter,
	successAdapter *json_success.Adapter,
	builder results.Builder,
) results.Adapter {
	out := Adapter{
		interruptionAdapter: interruptionAdapter,
		successAdapter:      successAdapter,
		builder:             builder,
	}

	return &out
}

// ToBytes converts instance to bytes
func (app *Adapter) ToBytes(ins results.Result) ([]byte, error) {
	ptr, err := app.ResultToStruct(ins)
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
func (app *Adapter) ToInstance(bytes []byte) (results.Result, error) {
	ins := new(Result)
	err := json.Unmarshal(bytes, ins)
	if err != nil {
		return nil, err
	}

	return app.StructToResult(*ins)
}

// ResultToStruct converts a result to struct
func (app *Adapter) ResultToStruct(ins results.Result) (*Result, error) {
	out := Result{}
	if ins.IsInterruption() {
		ptr, err := app.interruptionAdapter.InterruptionToStruct(ins.Interruption())
		if err != nil {
			return nil, err
		}

		out.Interruption = ptr
	}

	if ins.IsSuccess() {
		ptr, err := app.successAdapter.SuccessToStruct(ins.Success())
		if err != nil {
			return nil, err
		}

		out.Success = ptr
	}

	return &out, nil
}

// StructToResult converts a struct to result
func (app *Adapter) StructToResult(str Result) (results.Result, error) {
	builder := app.builder.Create()
	if str.Interruption != nil {
		ins, err := app.interruptionAdapter.StructToInterruption(*str.Interruption)
		if err != nil {
			return nil, err
		}

		builder.WithInterruption(ins)
	}

	if str.Success != nil {
		ins, err := app.successAdapter.StructToSuccess(*str.Success)
		if err != nil {
			return nil, err
		}

		builder.WithSuccess(ins)
	}

	return builder.Now()
}
