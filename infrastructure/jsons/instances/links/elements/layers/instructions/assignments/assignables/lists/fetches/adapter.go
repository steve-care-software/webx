package fetches

import (
	"encoding/json"

	"github.com/steve-care-software/datastencil/domain/instances/pointers/resources/logics/layers/instructions/assignments/assignables/lists/fetches"
)

// Adapter represents an adapter
type Adapter struct {
	builder fetches.Builder
}

func createAdapter(
	builder fetches.Builder,
) fetches.Adapter {
	out := Adapter{
		builder: builder,
	}

	return &out
}

// ToBytes converts instance to bytes
func (app *Adapter) ToBytes(ins fetches.Fetch) ([]byte, error) {
	ptr, err := app.FetchToStruct(ins)
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
func (app *Adapter) ToInstance(bytes []byte) (fetches.Fetch, error) {
	ins := new(Fetch)
	err := json.Unmarshal(bytes, ins)
	if err != nil {
		return nil, err
	}

	return app.StructToFetch(*ins)
}

// FetchToStruct converts a fetch to struct
func (app *Adapter) FetchToStruct(ins fetches.Fetch) (*Fetch, error) {
	return &Fetch{
		List:  ins.List(),
		Index: ins.Index(),
	}, nil
}

// StructToFetch converts a struct to fetch
func (app *Adapter) StructToFetch(str Fetch) (fetches.Fetch, error) {
	return app.builder.Create().
		WithList(str.List).
		WithIndex(str.Index).
		Now()
}
