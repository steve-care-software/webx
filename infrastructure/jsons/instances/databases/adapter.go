package databases

import (
	"encoding/json"

	"github.com/steve-care-software/datastencil/domain/instances/databases"
	json_commits "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/databases/commits"
)

// Adapter represents an adapter
type Adapter struct {
	commitAdapter *json_commits.Adapter
	builder       databases.Builder
}

func createAdapter(
	commitAdapter *json_commits.Adapter,
	builder databases.Builder,
) databases.Adapter {
	out := Adapter{
		commitAdapter: commitAdapter,
		builder:       builder,
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
	ptr, err := app.commitAdapter.CommitToStruct(ins.Head())
	if err != nil {
		return nil, err
	}

	out := Database{
		Path:        ins.Path(),
		Description: ins.Description(),
		Head:        *ptr,
		IsActive:    false,
	}

	if ins.IsActive() {
		out.IsActive = true
	}

	return &out, nil
}

// StructToCommit converts a struct to database
func (app *Adapter) StructToDatabase(str Database) (databases.Database, error) {
	ins, err := app.commitAdapter.StructToCommit(str.Head)
	if err != nil {
		return nil, err
	}

	builder := app.builder.Create().
		WithPath(str.Path).
		WithDescription(str.Description).
		WithHead(ins)

	if str.IsActive {
		builder.IsActive()
	}

	return builder.Now()
}
