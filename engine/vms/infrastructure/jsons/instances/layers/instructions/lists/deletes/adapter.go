package deletes

import (
	"bytes"
	"encoding/json"

	"github.com/steve-care-software/webx/engine/vms/domain/instances/layers/instructions/lists/deletes"
)

// Adapter represents a delete adapter
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
func (app *Adapter) ToInstance(data []byte) (deletes.Delete, error) {
	ins := new(Delete)
	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err := decoder.Decode(ins)
	if err != nil {
		return nil, err
	}

	return app.StructToDelete(*ins)
}

// DeleteToStruct converts a delete to struct
func (app *Adapter) DeleteToStruct(ins deletes.Delete) (*Delete, error) {
	return &Delete{
		List:  ins.List(),
		Index: ins.Idx(),
	}, nil
}

// StructToDelete converts a struct to delete
func (app *Adapter) StructToDelete(str Delete) (deletes.Delete, error) {
	return app.builder.Create().
		WithList(str.List).
		WithIndex(str.Index).
		Now()
}
