package inits

import (
	"bytes"
	"encoding/json"

	"github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables/executions/inits"
)

// Adapter represents the adapter
type Adapter struct {
	builder inits.Builder
}

func createAdapter(
	builder inits.Builder,
) inits.Adapter {
	out := Adapter{
		builder: builder,
	}

	return &out
}

// ToBytes converts an instance to bytes
func (app *Adapter) ToBytes(ins inits.Init) ([]byte, error) {
	str := app.InitToStruct(ins)
	js, err := json.Marshal(str)
	if err != nil {
		return nil, err
	}

	return js, nil
}

// ToInstance converts a bytes to instance
func (app *Adapter) ToInstance(data []byte) (inits.Init, error) {
	ins := new(Init)
	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err := decoder.Decode(ins)
	if err != nil {
		return nil, err
	}

	return app.StructToInit(*ins)
}

// InitToStruct converts an init to struct
func (app *Adapter) InitToStruct(ins inits.Init) Init {
	return Init{
		Path:        ins.Path(),
		Name:        ins.Name(),
		Description: ins.Description(),
		Context:     ins.Context(),
	}
}

// StructToInit converts a struct to init
func (app *Adapter) StructToInit(str Init) (inits.Init, error) {
	return app.builder.Create().
		WithPath(str.Path).
		WithName(str.Name).
		WithDescription(str.Description).
		WithContext(str.Context).
		Now()
}
