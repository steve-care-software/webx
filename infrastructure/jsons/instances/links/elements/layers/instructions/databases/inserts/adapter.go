package inserts

import (
	"encoding/json"

	"github.com/steve-care-software/datastencil/domain/instances/links/elements/layers/instructions/databases/inserts"
)

// Adapter represents an adapter
type Adapter struct {
	builder inserts.Builder
}

func createAdapter(
	builder inserts.Builder,
) inserts.Adapter {
	out := Adapter{
		builder: builder,
	}

	return &out
}

// ToBytes converts an instance to bytes
func (app *Adapter) ToBytes(ins inserts.Insert) ([]byte, error) {
	ptr, err := app.InsertToStruct(ins)
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
func (app *Adapter) ToInstance(bytes []byte) (inserts.Insert, error) {
	ins := new(Insert)
	err := json.Unmarshal(bytes, ins)
	if err != nil {
		return nil, err
	}

	return app.StructToInsert(*ins)
}

// InsertToStruct converts a insert to struct
func (app *Adapter) InsertToStruct(ins inserts.Insert) (*Insert, error) {
	return &Insert{
		Context:  ins.Context(),
		Path:     ins.Path(),
		Instance: ins.Instance(),
	}, nil
}

// StructToInsert converts a struct to insert
func (app *Adapter) StructToInsert(str Insert) (inserts.Insert, error) {
	return app.builder.Create().
		WithContext(str.Context).
		WithPath(str.Path).
		WithInstance(str.Instance).
		Now()
}
