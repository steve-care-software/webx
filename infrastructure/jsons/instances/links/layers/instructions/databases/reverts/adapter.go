package reverts

import (
	"encoding/json"

	"github.com/steve-care-software/datastencil/domain/instances/links/layers/instructions/databases/reverts"
)

// Adapter represents an adapter
type Adapter struct {
	builder reverts.Builder
}

func createAdapter(
	builder reverts.Builder,
) reverts.Adapter {
	out := Adapter{
		builder: builder,
	}

	return &out
}

// ToBytes converts instance to bytes
func (app *Adapter) ToBytes(ins reverts.Revert) ([]byte, error) {
	ptr, err := app.RevertToStruct(ins)
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
func (app *Adapter) ToInstance(bytes []byte) (reverts.Revert, error) {
	ins := new(Revert)
	err := json.Unmarshal(bytes, ins)
	if err != nil {
		return nil, err
	}

	return app.StructToInsert(*ins)
}

// RevertToStruct converts a revert to struct
func (app *Adapter) RevertToStruct(ins reverts.Revert) (*Revert, error) {
	out := Revert{}
	if ins.HasIndex() {
		out.Index = ins.Index()
	}

	return &out, nil
}

// StructToInsert converts a struct to insert
func (app *Adapter) StructToInsert(str Revert) (reverts.Revert, error) {
	builder := app.builder.Create()
	if str.Index != "" {
		builder.WithIndex(str.Index)
	}

	return builder.Now()
}
