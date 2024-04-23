package inserts

import (
	"encoding/json"

	"github.com/steve-care-software/datastencil/domain/instances/links/layers/instructions/accounts/inserts"
)

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
	str := app.InsertToStruct(ins)
	js, err := json.Marshal(str)
	if err != nil {
		return nil, err
	}

	return js, nil
}

// ToInstance converts bytes to an instance
func (app *Adapter) ToInstance(bytes []byte) (inserts.Insert, error) {
	ins := new(Insert)
	err := json.Unmarshal(bytes, ins)
	if err != nil {
		return nil, err
	}

	return app.StructToInsert(*ins)
}

// InsertToStruct converts an insert to struct
func (app *Adapter) InsertToStruct(ins inserts.Insert) Insert {
	return Insert{
		Username: ins.Username(),
		Password: ins.Password(),
	}
}

// StructToInsert converts a struct to insert
func (app *Adapter) StructToInsert(str Insert) (inserts.Insert, error) {
	return app.builder.Create().
		WithUsername(str.Username).
		WithPassword(str.Password).
		Now()
}
