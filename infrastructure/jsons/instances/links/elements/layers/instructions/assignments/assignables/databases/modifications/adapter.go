package modifications

import (
	"encoding/json"

	"github.com/steve-care-software/datastencil/domain/instances/links/elements/layers/instructions/assignments/assignables/databases/modifications"
)

// Adapter represents an adapter
type Adapter struct {
	builder modifications.Builder
}

func createAdapter(
	builder modifications.Builder,
) modifications.Adapter {
	out := Adapter{
		builder: builder,
	}

	return &out
}

// ToBytes converts an instance to bytes
func (app *Adapter) ToBytes(ins modifications.Modification) ([]byte, error) {
	ptr, err := app.ModificationToStruct(ins)
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
func (app *Adapter) ToInstance(bytes []byte) (modifications.Modification, error) {
	ins := new(Modification)
	err := json.Unmarshal(bytes, ins)
	if err != nil {
		return nil, err
	}

	return app.StructToModification(*ins)
}

// ModificationToStruct converts a modification to struct
func (app *Adapter) ModificationToStruct(ins modifications.Modification) (*Modification, error) {
	out := Modification{}
	if ins.IsInsert() {
		out.Insert = ins.Insert()
	}

	if ins.IsDelete() {
		out.Delete = ins.Delete()
	}

	return &out, nil
}

// StructToModification converts a struct to modification
func (app *Adapter) StructToModification(str Modification) (modifications.Modification, error) {
	builder := app.builder.Create()
	if str.Insert != "" {
		builder.WithInsert(str.Insert)
	}

	if str.Delete != "" {
		builder.WithDelete(str.Delete)
	}

	return builder.Now()
}
