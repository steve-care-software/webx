package databases

import (
	"encoding/json"

	"github.com/steve-care-software/datastencil/domain/instances/pointers/resources/logics/bridges/layers/instructions/databases"
)

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

// ToBytes converts an instance to bytes
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

// ToInstance converts bytes to an instance
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
	out := Database{}
	if ins.IsSave() {
		out.Save = ins.Save()
	}

	if ins.IsDelete() {
		out.Delete = ins.Delete()
	}

	return &out, nil
}

// StructToDatabase converts a struct to database
func (app *Adapter) StructToDatabase(str Database) (databases.Database, error) {
	builder := app.builder.Create()
	if str.Save != "" {
		builder.WithSave(str.Save)
	}

	if str.Delete != "" {
		builder.WithDelete(str.Delete)
	}

	return builder.Now()
}
