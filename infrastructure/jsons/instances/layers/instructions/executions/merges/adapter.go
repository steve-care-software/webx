package merges

import (
	"bytes"
	"encoding/json"

	"github.com/steve-care-software/datastencil/domain/instances/layers/instructions/executions/merges"
)

// Adapter represents a merge adapter
type Adapter struct {
	builder merges.Builder
}

func createAdapter(
	builder merges.Builder,
) merges.Adapter {
	out := Adapter{
		builder: builder,
	}

	return &out
}

// ToBytes converts instance to bytes
func (app *Adapter) ToBytes(ins merges.Merge) ([]byte, error) {
	str := app.MergeToStruct(ins)
	js, err := json.Marshal(str)
	if err != nil {
		return nil, err
	}

	return js, nil
}

// ToInstance converts bytes to instance
func (app *Adapter) ToInstance(data []byte) (merges.Merge, error) {
	ins := new(Merge)
	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err := decoder.Decode(ins)
	if err != nil {
		return nil, err
	}

	return app.StructToMerge(*ins)
}

// MergeToStruct converts a merge to struct
func (app *Adapter) MergeToStruct(ins merges.Merge) Merge {
	return Merge{
		Base: ins.Base(),
		Top:  ins.Top(),
	}
}

// StructToMerge converts a struct to merge
func (app *Adapter) StructToMerge(str Merge) (merges.Merge, error) {
	return app.builder.Create().
		WithBase(str.Base).
		WithTop(str.Top).
		Now()
}
