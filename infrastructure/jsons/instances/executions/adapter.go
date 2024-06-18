package executions

import (
	"encoding/json"

	"github.com/steve-care-software/datastencil/domain/instances/executions"
	json_databases "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/databases"
	json_links "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/executions/links"
)

// Adapter represents an adapter
type Adapter struct {
	databaseAdapter  *json_databases.Adapter
	linkAdapter      *json_links.Adapter
	builder          executions.Builder
	executionBuilder executions.ExecutionBuilder
}

func createAdapter(
	databaseAdapter *json_databases.Adapter,
	linkAdapter *json_links.Adapter,
	builder executions.Builder,
	executionBuilder executions.ExecutionBuilder,
) executions.Adapter {
	out := Adapter{
		databaseAdapter:  databaseAdapter,
		linkAdapter:      linkAdapter,
		builder:          builder,
		executionBuilder: executionBuilder,
	}

	return &out
}

// ToBytes converts instance to bytes
func (app *Adapter) ToBytes(ins executions.Executions) ([]byte, error) {
	ptr, err := app.ExecutionsToStruct(ins)
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
func (app *Adapter) ToInstance(bytes []byte) (executions.Executions, error) {
	ins := new([]Execution)
	err := json.Unmarshal(bytes, ins)
	if err != nil {
		return nil, err
	}

	return app.StructToExecutions(*ins)
}

// ExecutionsToStruct converts an executions to struct
func (app *Adapter) ExecutionsToStruct(ins executions.Executions) ([]Execution, error) {
	out := []Execution{}
	list := ins.List()
	for _, oneIns := range list {
		ptr, err := app.ExecutionToStruct(oneIns)
		if err != nil {
			return nil, err
		}

		out = append(out, *ptr)
	}

	return out, nil
}

// StructToExecutions converts a struct to executions
func (app *Adapter) StructToExecutions(str []Execution) (executions.Executions, error) {
	list := []executions.Execution{}
	for _, oneStruct := range str {
		ins, err := app.StructToExecution(oneStruct)
		if err != nil {
			return nil, err
		}

		list = append(list, ins)
	}

	return app.builder.Create().
		WithList(list).
		Now()
}

// ExecutionToStruct converts an execution to struct
func (app *Adapter) ExecutionToStruct(ins executions.Execution) (*Execution, error) {
	ptrLink, err := app.linkAdapter.LinkToStruct(ins.Logic())
	if err != nil {
		return nil, err
	}

	ptrDatabase, err := app.databaseAdapter.DatabaseToStruct(ins.Database())
	if err != nil {
		return nil, err
	}

	return &Execution{
		Link:     *ptrLink,
		Database: *ptrDatabase,
	}, nil
}

// StructToExecution converts a struct to execution
func (app *Adapter) StructToExecution(str Execution) (executions.Execution, error) {
	linkIns, err := app.linkAdapter.StructToLink(str.Link)
	if err != nil {
		return nil, err
	}

	databaseIns, err := app.databaseAdapter.StructToDatabase(str.Database)
	if err != nil {
		return nil, err
	}

	return app.executionBuilder.Create().
		WithDatabase(databaseIns).
		WithLogic(linkIns).
		Now()
}
