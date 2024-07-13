package failures

import (
	"encoding/json"

	"github.com/steve-care-software/datastencil/domain/instances/executions/layers/results/interruptions/failures"
)

// Adapter represents an adapter
type Adapter struct {
	builder failures.Builder
}

func createAdapter(
	builder failures.Builder,
) failures.Adapter {
	out := Adapter{
		builder: builder,
	}

	return &out
}

// ToBytes converts instance to bytes
func (app *Adapter) ToBytes(ins failures.Failure) ([]byte, error) {
	ptr, err := app.FailureToStruct(ins)
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
func (app *Adapter) ToInstance(bytes []byte) (failures.Failure, error) {
	ins := new(Failure)
	err := json.Unmarshal(bytes, ins)
	if err != nil {
		return nil, err
	}

	return app.StructToFailure(*ins)
}

// FailureToStruct converts a failure to struct
func (app *Adapter) FailureToStruct(ins failures.Failure) (*Failure, error) {
	return &Failure{
		Index:           ins.Index(),
		Code:            ins.Code(),
		Message:         ins.Message(),
		IsRaisedInLayer: ins.IsRaisedInLayer(),
	}, nil
}

// StructToFailure converts a struct to failure
func (app *Adapter) StructToFailure(str Failure) (failures.Failure, error) {
	builder := app.builder.Create().
		WithIndex(str.Index).
		WithCode(str.Code).
		WithMessage(str.Message)

	if str.IsRaisedInLayer {
		builder.IsRaisedInLayer()
	}

	return builder.Now()
}
