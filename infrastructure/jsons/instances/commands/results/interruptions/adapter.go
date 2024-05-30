package interruptions

import (
	"encoding/json"

	"github.com/steve-care-software/datastencil/domain/instances/executions/links/layers/results/interruptions"
	json_failures "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/commands/results/interruptions/failures"
)

// Adapter represents an adapter
type Adapter struct {
	failureAdapter *json_failures.Adapter
	builder        interruptions.Builder
}

func createAdapter(
	failureAdapter *json_failures.Adapter,
	builder interruptions.Builder,
) interruptions.Adapter {
	out := Adapter{
		failureAdapter: failureAdapter,
		builder:        builder,
	}

	return &out
}

// ToBytes converts instance to bytes
func (app *Adapter) ToBytes(ins interruptions.Interruption) ([]byte, error) {
	ptr, err := app.InterruptionToStruct(ins)
	if err != nil {
		return nil, err
	}

	js, err := json.Marshal(ptr)
	if err != nil {
		return nil, err
	}

	return js, nil
}

// ToInstance converts bytes to interruption
func (app *Adapter) ToInstance(bytes []byte) (interruptions.Interruption, error) {
	ins := new(Interruption)
	err := json.Unmarshal(bytes, ins)
	if err != nil {
		return nil, err
	}

	return app.StructToInterruption(*ins)
}

// InterruptionToStruct converts an interruption to struct
func (app *Adapter) InterruptionToStruct(ins interruptions.Interruption) (*Interruption, error) {
	out := Interruption{}
	if ins.IsStop() {
		out.Stop = ins.Stop()
	}

	if ins.IsFailure() {
		ptr, err := app.failureAdapter.FailureToStruct(ins.Failure())
		if err != nil {
			return nil, err
		}

		out.Failure = ptr
	}

	return &out, nil
}

// StructToFailure converts a struct to failure
func (app *Adapter) StructToInterruption(str Interruption) (interruptions.Interruption, error) {
	builder := app.builder.Create()
	if str.Stop != nil {
		builder.WithStop(*str.Stop)
	}

	if str.Failure != nil {
		ins, err := app.failureAdapter.StructToFailure(*str.Failure)
		if err != nil {
			return nil, err
		}

		builder.WithFailure(ins)
	}

	return builder.Now()
}
