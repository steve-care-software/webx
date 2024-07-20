package lists

import (
	"bytes"
	"encoding/json"

	"github.com/steve-care-software/webx/engine/stencils/domain/instances/layers/instructions/lists"
	json_deletes "github.com/steve-care-software/webx/engine/stencils/infrastructure/jsons/instances/layers/instructions/lists/deletes"
	json_inserts "github.com/steve-care-software/webx/engine/stencils/infrastructure/jsons/instances/layers/instructions/lists/inserts"
)

// Adapter represents the list adapter
type Adapter struct {
	deleteAdapter *json_deletes.Adapter
	insertAdapter *json_inserts.Adapter
	builder       lists.Builder
}

func createAdapter(
	deleteAdapter *json_deletes.Adapter,
	insertAdapter *json_inserts.Adapter,
	builder lists.Builder,
) lists.Adapter {
	out := Adapter{
		deleteAdapter: deleteAdapter,
		insertAdapter: insertAdapter,
		builder:       builder,
	}

	return &out
}

// ToBytes converts instance to bytes
func (app *Adapter) ToBytes(ins lists.List) ([]byte, error) {
	ptr, err := app.ListToStruct(ins)
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
func (app *Adapter) ToInstance(data []byte) (lists.List, error) {
	ins := new(List)
	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err := decoder.Decode(ins)
	if err != nil {
		return nil, err
	}

	return app.StructToList(*ins)
}

// ListToStruct converts a list to struct
func (app *Adapter) ListToStruct(ins lists.List) (*List, error) {
	out := List{}
	if ins.IsInsert() {
		ptr, err := app.insertAdapter.InsertToStruct(ins.Insert())
		if err != nil {
			return nil, err
		}

		out.Insert = ptr
	}

	if ins.IsDelete() {
		ptr, err := app.deleteAdapter.DeleteToStruct(ins.Delete())
		if err != nil {
			return nil, err
		}

		out.Delete = ptr
	}

	return &out, nil
}

// StructToList converts a struct to list
func (app *Adapter) StructToList(str List) (lists.List, error) {
	builder := app.builder.Create()
	if str.Insert != nil {
		ins, err := app.insertAdapter.StructToInsert(*str.Insert)
		if err != nil {
			return nil, err
		}

		builder.WithInsert(ins)
	}

	if str.Delete != nil {
		ins, err := app.deleteAdapter.StructToDelete(*str.Delete)
		if err != nil {
			return nil, err
		}

		builder.WithDelete(ins)
	}

	return builder.Now()
}
