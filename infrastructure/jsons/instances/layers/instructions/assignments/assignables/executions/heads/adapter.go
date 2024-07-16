package heads

import (
	"bytes"
	"encoding/json"

	"github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables/executions/heads"
)

// Adapter represents the adapter
type Adapter struct {
	builder heads.Builder
}

func createAdapter(
	builder heads.Builder,
) heads.Adapter {
	out := Adapter{
		builder: builder,
	}

	return &out
}

// ToBytes converts an instance to bytes
func (app *Adapter) ToBytes(ins heads.Head) ([]byte, error) {
	str := app.HeadToStruct(ins)
	js, err := json.Marshal(str)
	if err != nil {
		return nil, err
	}

	return js, nil
}

// ToInstance converts a bytes to instance
func (app *Adapter) ToInstance(data []byte) (heads.Head, error) {
	ins := new(Head)
	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err := decoder.Decode(ins)
	if err != nil {
		return nil, err
	}

	return app.StructToHead(*ins)
}

// HeadToStruct converts an head to struct
func (app *Adapter) HeadToStruct(ins heads.Head) Head {
	return Head{
		Context: ins.Context(),
		Return:  ins.ReturnHash(),
	}
}

// StructToHead converts a struct to input
func (app *Adapter) StructToHead(str Head) (heads.Head, error) {
	return app.builder.Create().
		WithContext(str.Context).
		WithReturn(str.Return).
		Now()
}
