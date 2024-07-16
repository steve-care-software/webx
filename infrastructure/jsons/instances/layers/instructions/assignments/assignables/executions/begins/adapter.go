package begins

import (
	"bytes"
	"encoding/json"

	"github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables/executions/begins"
)

// Adapter represents the begin adapter
type Adapter struct {
	builder begins.Builder
}

func createAdapter(
	builder begins.Builder,
) begins.Adapter {
	out := Adapter{
		builder: builder,
	}

	return &out
}

// ToBytes converts an instance to bytes
func (app *Adapter) ToBytes(ins begins.Begin) ([]byte, error) {
	str := app.BeginToStruct(ins)
	js, err := json.Marshal(str)
	if err != nil {
		return nil, err
	}

	return js, nil
}

// ToInstance converts a bytes to instance
func (app *Adapter) ToInstance(data []byte) (begins.Begin, error) {
	ins := new(Begin)
	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err := decoder.Decode(ins)
	if err != nil {
		return nil, err
	}

	return app.StructToBegin(*ins)
}

// BeginToStruct converts a begin to struct
func (app *Adapter) BeginToStruct(ins begins.Begin) Begin {
	return Begin{
		Context: ins.Context(),
		Path:    ins.Path(),
	}
}

// StructToBegin converts a struct to begin
func (app *Adapter) StructToBegin(str Begin) (begins.Begin, error) {
	return app.builder.Create().
		WithContext(str.Context).
		WithPath(str.Path).
		Now()
}
