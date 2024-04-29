package deletes

import (
	"encoding/json"

	"github.com/steve-care-software/datastencil/domain/instances/links/layers/instructions/databases/deletes"
)

// Adapter represents an adapter
type Adapter struct {
	builder deletes.Builder
}

func createAdapter(
	builder deletes.Builder,
) deletes.Adapter {
	out := Adapter{
		builder: builder,
	}

	return &out
}

// ToBytes converts an instance to bytes
func (app *Adapter) ToBytes(ins deletes.Delete) ([]byte, error) {
	ptr, err := app.DeleteToStruct(ins)
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
func (app *Adapter) ToInstance(bytes []byte) (deletes.Delete, error) {
	ins := new(Delete)
	err := json.Unmarshal(bytes, ins)
	if err != nil {
		return nil, err
	}

	return app.StructToDelete(*ins)
}

// DeleteToStruct converts a delete to struct
func (app *Adapter) DeleteToStruct(ins deletes.Delete) (*Delete, error) {
	return &Delete{
		Context:    ins.Context(),
		Path:       ins.Path(),
		Identifier: ins.Identifier(),
	}, nil
}

// StructToDelete converts a struct to delete
func (app *Adapter) StructToDelete(str Delete) (deletes.Delete, error) {
	return app.builder.Create().
		WithContext(str.Context).
		WithPath(str.Path).
		WithIdentifier(str.Identifier).
		Now()
}
