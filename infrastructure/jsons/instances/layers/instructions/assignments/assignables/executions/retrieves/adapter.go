package retrieves

import (
	"encoding/json"

	"github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables/executions/retrieves"
)

// Adapter represents the adapter
type Adapter struct {
	builder retrieves.Builder
}

func createAdapter(
	builder retrieves.Builder,
) retrieves.Adapter {
	out := Adapter{
		builder: builder,
	}

	return &out
}

// ToBytes converts an instance to bytes
func (app *Adapter) ToBytes(ins retrieves.Retrieve) ([]byte, error) {
	str := app.RetrieveToStruct(ins)
	js, err := json.Marshal(str)
	if err != nil {
		return nil, err
	}

	return js, nil
}

// ToInstance converts a bytes to instance
func (app *Adapter) ToInstance(bytes []byte) (retrieves.Retrieve, error) {
	ins := new(Retrieve)
	err := json.Unmarshal(bytes, ins)
	if err != nil {
		return nil, err
	}

	return app.StructToRetrieve(*ins)
}

// RetrieveToStruct converts a retrieve to struct
func (app *Adapter) RetrieveToStruct(ins retrieves.Retrieve) Retrieve {
	out := Retrieve{
		Context: ins.Context(),
		Index:   ins.Index(),
		Return:  ins.Return(),
	}

	if ins.HasLength() {
		out.Length = ins.Length()
	}

	return out
}

// StructToRetrieve converts a struct to retrieve
func (app *Adapter) StructToRetrieve(str Retrieve) (retrieves.Retrieve, error) {
	builder := app.builder.Create().
		WithContext(str.Context).
		WithIndex(str.Index).
		WithReturn(str.Return)

	if str.Length != "" {
		builder.WithLength(str.Length)
	}

	return builder.Now()
}
