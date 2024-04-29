package databases

import (
	"encoding/json"

	"github.com/steve-care-software/datastencil/domain/instances/links/layers/instructions/databases"
	json_deletes "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/links/layers/instructions/databases/deletes"
	json_inserts "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/links/layers/instructions/databases/inserts"
	json_reverts "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/links/layers/instructions/databases/reverts"
)

// Adapter represents an adapter
type Adapter struct {
	insertAdapter *json_inserts.Adapter
	deleteAdapter *json_deletes.Adapter
	revertAdapter *json_reverts.Adapter
	builder       databases.Builder
}

func createAdapter(
	insertAdapter *json_inserts.Adapter,
	deleteAdapter *json_deletes.Adapter,
	revertAdapter *json_reverts.Adapter,
	builder databases.Builder,
) databases.Adapter {
	out := Adapter{
		insertAdapter: insertAdapter,
		deleteAdapter: deleteAdapter,
		revertAdapter: revertAdapter,
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
	out := Database{}
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

	if ins.IsCommit() {
		out.Commit = ins.Commit()
	}

	if ins.IsCancel() {
		out.Cancel = ins.Cancel()
	}

	if ins.IsRevert() {
		ptr, err := app.revertAdapter.RevertToStruct(ins.Revert())
		if err != nil {
			return nil, err
		}

		out.Revert = ptr
	}

	return &out, nil
}

// StructToDatabase converts a struct to database
func (app *Adapter) StructToDatabase(str Database) (databases.Database, error) {
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

	if str.Commit != "" {
		builder.WithCommit(str.Commit)
	}

	if str.Cancel != "" {
		builder.WithCancel(str.Cancel)
	}

	if str.Revert != nil {
		ins, err := app.revertAdapter.StructToRevert(*str.Revert)
		if err != nil {
			return nil, err
		}

		builder.WithRevert(ins)
	}

	return builder.Now()
}
