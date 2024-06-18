package databases

import (
	"encoding/json"

	"github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables/databases"
	json_actions "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/layers/instructions/assignments/assignables/databases/actions"
	json_commits "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/layers/instructions/assignments/assignables/databases/commits"
	json_databases "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/layers/instructions/assignments/assignables/databases/databases"
	json_deletes "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/layers/instructions/assignments/assignables/databases/deletes"
	json_modifications "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/layers/instructions/assignments/assignables/databases/modifications"
	json_retrieves "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/layers/instructions/assignments/assignables/databases/retrieves"
)

// Adapter represents an adapter
type Adapter struct {
	actionAdapter       *json_actions.Adapter
	commitAdapter       *json_commits.Adapter
	databaseAdapter     *json_databases.Adapter
	deleteAdapter       *json_deletes.Adapter
	modificationAdapter *json_modifications.Adapter
	retrieveAdapter     *json_retrieves.Adapter
	builder             databases.Builder
}

func createAdapter(
	actionAdapter *json_actions.Adapter,
	commitAdapter *json_commits.Adapter,
	databaseAdapter *json_databases.Adapter,
	deleteAdapter *json_deletes.Adapter,
	modificationAdapter *json_modifications.Adapter,
	retrieveAdapter *json_retrieves.Adapter,
	builder databases.Builder,
) databases.Adapter {
	out := Adapter{
		actionAdapter:       actionAdapter,
		commitAdapter:       commitAdapter,
		databaseAdapter:     databaseAdapter,
		deleteAdapter:       deleteAdapter,
		modificationAdapter: modificationAdapter,
		retrieveAdapter:     retrieveAdapter,
		builder:             builder,
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
	if ins.IsAction() {
		ptr, err := app.actionAdapter.ActionToStruct(ins.Action())
		if err != nil {
			return nil, err
		}

		out.Action = ptr
	}

	if ins.IsCommit() {
		ptr, err := app.commitAdapter.CommitToStruct(ins.Commit())
		if err != nil {
			return nil, err
		}

		out.Commit = ptr
	}

	if ins.IsDatabase() {
		ptr, err := app.databaseAdapter.DatabaseToStruct(ins.Database())
		if err != nil {
			return nil, err
		}

		out.Database = ptr
	}

	if ins.IsDelete() {
		ptr, err := app.deleteAdapter.DeleteToStruct(ins.Delete())
		if err != nil {
			return nil, err
		}

		out.Delete = ptr
	}

	if ins.IsModification() {
		ptr, err := app.modificationAdapter.ModificationToStruct(ins.Modification())
		if err != nil {
			return nil, err
		}

		out.Modification = ptr
	}

	if ins.IsRetrieve() {
		ptr, err := app.retrieveAdapter.RetrieveToStruct(ins.Retrieve())
		if err != nil {
			return nil, err
		}

		out.Retrieve = ptr
	}

	return &out, nil
}

// StructToDatabase converts a struct to database
func (app *Adapter) StructToDatabase(str Database) (databases.Database, error) {
	builder := app.builder.Create()
	if str.Action != nil {
		ins, err := app.actionAdapter.StructToAction(*str.Action)
		if err != nil {
			return nil, err
		}

		builder.WithAction(ins)
	}

	if str.Commit != nil {
		ins, err := app.commitAdapter.StructToCommit(*str.Commit)
		if err != nil {
			return nil, err
		}

		builder.WithCommit(ins)
	}

	if str.Database != nil {
		ins, err := app.databaseAdapter.StructToDatabase(*str.Database)
		if err != nil {
			return nil, err
		}

		builder.WithDatabase(ins)
	}

	if str.Delete != nil {
		ins, err := app.deleteAdapter.StructToDelete(*str.Delete)
		if err != nil {
			return nil, err
		}

		builder.WithDelete(ins)
	}

	if str.Modification != nil {
		ins, err := app.modificationAdapter.StructToModification(*str.Modification)
		if err != nil {
			return nil, err
		}

		builder.WithModification(ins)
	}

	if str.Retrieve != nil {
		ins, err := app.retrieveAdapter.StructToRetrieve(*str.Retrieve)
		if err != nil {
			return nil, err
		}

		builder.WithRetrieve(ins)
	}

	return builder.Now()
}
