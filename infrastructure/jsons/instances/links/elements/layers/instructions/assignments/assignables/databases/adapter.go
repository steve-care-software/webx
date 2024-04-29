package databases

import (
	"encoding/json"

	"github.com/steve-care-software/datastencil/domain/instances/links/elements/layers/instructions/assignments/assignables/databases"
	json_repositories "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/links/elements/layers/instructions/assignments/assignables/databases/repositories"
	json_services "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/links/elements/layers/instructions/assignments/assignables/databases/services"
)

// Adapter represents an adapter
type Adapter struct {
	repositoryAdapter *json_repositories.Adapter
	serviceAdapter    *json_services.Adapter
	builder           databases.Builder
}

func createAdapter(
	repositoryAdapter *json_repositories.Adapter,
	serviceAdapter *json_services.Adapter,
	builder databases.Builder,
) databases.Adapter {
	out := Adapter{
		repositoryAdapter: repositoryAdapter,
		serviceAdapter:    serviceAdapter,
		builder:           builder,
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

// DatabaseToStruct converts a databasse to struct
func (app *Adapter) DatabaseToStruct(ins databases.Database) (*Database, error) {
	out := Database{}
	if ins.IsRepository() {
		ptr, err := app.repositoryAdapter.RepositoryToStruct(ins.Repository())
		if err != nil {
			return nil, err
		}

		out.Repository = ptr
	}

	if ins.IsService() {
		ptr, err := app.serviceAdapter.ServiceToStruct(ins.Service())
		if err != nil {
			return nil, err
		}

		out.Service = ptr
	}

	return &out, nil
}

// StructToDatabase converts a struct to database
func (app *Adapter) StructToDatabase(str Database) (databases.Database, error) {
	builder := app.builder.Create()
	if str.Repository != nil {
		ins, err := app.repositoryAdapter.StructToRepository(*str.Repository)
		if err != nil {
			return nil, err
		}

		builder.WithRepository(ins)
	}

	if str.Service != nil {
		ins, err := app.serviceAdapter.StructToService(*str.Service)
		if err != nil {
			return nil, err
		}

		builder.WithService(ins)
	}

	return builder.Now()
}
