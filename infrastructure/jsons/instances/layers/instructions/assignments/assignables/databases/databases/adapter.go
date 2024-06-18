package databases

import (
	"encoding/json"

	"github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables/databases/databases"
)

// Adapter represents an adapter
type Adapter struct {
	builder databases.Builder
}

func createAdapter(
	builder databases.Builder,
) databases.Adapter {
	out := Adapter{
		builder: builder,
	}

	return &out
}

// ToBytes converts instance to bytes
func (app *Adapter) ToBytes(ins databases.Database) ([]byte, error) {
	ptr, err := app.DatabaseToStruct(ins)
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
func (app *Adapter) ToInstance(bytes []byte) (databases.Database, error) {
	ins := new(Database)
	err := json.Unmarshal(bytes, ins)
	if err != nil {
		return nil, err
	}

	return app.StructToDatabase(*ins)
}

// DatabaseToStruct converts a database to struct
func (app *Adapter) DatabaseToStruct(ins databases.Database) (*Database, error) {
	return &Database{
		Path:        ins.Path(),
		Description: ins.Description(),
		Head:        ins.Head(),
		IsActive:    ins.IsActive(),
	}, nil
}

// StructToDatabase converts a struct to database
func (app *Adapter) StructToDatabase(str Database) (databases.Database, error) {
	return app.builder.Create().
		WithPath(str.Path).
		WithDescription(str.Description).
		WithHead(str.Head).
		WithActive(str.IsActive).
		Now()
}
