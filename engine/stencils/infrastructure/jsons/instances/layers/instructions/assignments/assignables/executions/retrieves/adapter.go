package retrieves

import (
	"bytes"
	"encoding/json"

	"github.com/steve-care-software/webx/engine/stencils/domain/instances/layers/instructions/assignments/assignables/executions/retrieves"
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
func (app *Adapter) ToInstance(data []byte) (retrieves.Retrieve, error) {
	ins := new(Retrieve)
	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err := decoder.Decode(ins)
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
		WithIndex(str.Index)

	if str.Length != "" {
		builder.WithLength(str.Length)
	}

	return builder.Now()
}
