package deletes

import (
	"encoding/json"

	"github.com/steve-care-software/datastencil/domain/instances/databases/commits/actions/modifications/deletes"
)

// Adapter creates a new adapter
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

// ToBytes converts instance to bytes
func (app *Adapter) ToBytes(ins deletes.Delete) ([]byte, error) {
	str := app.DeleteToStruct(ins)
	js, err := json.Marshal(str)
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
func (app *Adapter) DeleteToStruct(ins deletes.Delete) Delete {
	return Delete{
		Index:  ins.Index(),
		Length: ins.Length(),
	}
}

// StructToDelete converts a struct to delete
func (app *Adapter) StructToDelete(str Delete) (deletes.Delete, error) {
	return app.builder.Create().
		WithIndex(str.Index).
		WithLength(str.Length).
		Now()
}