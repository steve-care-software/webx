package inserts

import (
	"bytes"
	"encoding/json"

	"github.com/steve-care-software/webx/engine/vms/domain/instances/layers/instructions/lists/inserts"
)

// Adapter represents an insert adapter
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
func (app *Adapter) ToInstance(data []byte) (inserts.Insert, error) {
	ins := new(Insert)
	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err := decoder.Decode(ins)
	if err != nil {
		return nil, err
	}

	return app.StructToInsert(*ins)
}

// InsertToStruct converts an insert to struct
func (app *Adapter) InsertToStruct(ins inserts.Insert) (*Insert, error) {
	return &Insert{
		List:    ins.List(),
		Element: ins.Element(),
	}, nil
}

// StructToInsert converts a struct to insert
func (app *Adapter) StructToInsert(str Insert) (inserts.Insert, error) {
	return app.builder.Create().
		WithList(str.List).
		WithElement(str.Element).
		Now()
}
