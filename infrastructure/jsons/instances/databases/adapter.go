package databases

import (
	"encoding/json"

	"github.com/steve-care-software/datastencil/domain/instances/databases"
	json_commits "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/databases/commits"
	json_heads "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/databases/heads"
)

// Adapter represents an adapter
type Adapter struct {
	commitAdapter *json_commits.Adapter
	headAdapter   *json_heads.Adapter
	builder       databases.Builder
}

func createAdapter(
	commitAdapter *json_commits.Adapter,
	headAdapter *json_heads.Adapter,
	builder databases.Builder,
) databases.Adapter {
	out := Adapter{
		commitAdapter: commitAdapter,
		headAdapter:   headAdapter,
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
	commitPtr, err := app.commitAdapter.CommitToStruct(ins.Commit())
	if err != nil {
		return nil, err
	}

	headPtr, err := app.headAdapter.HeadToStruct(ins.Head())
	if err != nil {
		return nil, err
	}

	return &Database{
		Commit: *commitPtr,
		Head:   *headPtr,
	}, nil
}

// StructToCommit converts a struct to database
func (app *Adapter) StructToDatabase(str Database) (databases.Database, error) {
	commitIns, err := app.commitAdapter.StructToCommit(str.Commit)
	if err != nil {
		return nil, err
	}

	headIns, err := app.headAdapter.StructToHead(str.Head)
	if err != nil {
		return nil, err
	}

	return app.builder.Create().
		WithCommit(commitIns).
		WithHead(headIns).
		Now()
}
