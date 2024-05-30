package lists

import (
	"encoding/json"

	"github.com/steve-care-software/datastencil/domain/instances/pointers/resources/logics/bridges/layers/instructions/assignments/assignables/lists"
	json_fetches "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/links/elements/layers/instructions/assignments/assignables/lists/fetches"
)

// Adapter represents an adapter
type Adapter struct {
	fetchAdapter *json_fetches.Adapter
	builder      lists.Builder
}

func createAdapter(
	fetchAdapter *json_fetches.Adapter,
	builder lists.Builder,
) lists.Adapter {
	out := Adapter{
		fetchAdapter: fetchAdapter,
		builder:      builder,
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
func (app *Adapter) ToInstance(bytes []byte) (lists.List, error) {
	ins := new(List)
	err := json.Unmarshal(bytes, ins)
	if err != nil {
		return nil, err
	}

	return app.StructToList(*ins)
}

// ListToStruct converts a list to struct
func (app *Adapter) ListToStruct(ins lists.List) (*List, error) {
	out := List{}
	if ins.IsFetch() {
		ptr, err := app.fetchAdapter.FetchToStruct(ins.Fetch())
		if err != nil {
			return nil, err
		}

		out.Fetch = ptr
	}

	if ins.IsCreate() {
		out.Create = ins.Create()
	}

	if ins.IsLength() {
		out.Length = ins.Length()
	}

	return &out, nil
}

// StructToList converts a struct to list
func (app *Adapter) StructToList(str List) (lists.List, error) {
	builder := app.builder.Create()
	if str.Fetch != nil {
		ins, err := app.fetchAdapter.StructToFetch(*str.Fetch)
		if err != nil {
			return nil, err
		}

		builder.WithFetch(ins)
	}

	if str.Length != "" {
		builder.WithLength(str.Length)
	}

	if str.Create != "" {
		builder.WithCreate(str.Create)
	}

	return builder.Now()
}
