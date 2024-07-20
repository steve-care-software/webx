package executions

import (
	"bytes"
	"encoding/json"

	"github.com/steve-care-software/webx/engine/stencils/domain/instances/layers/instructions/assignments/assignables/executions"
	json_executes "github.com/steve-care-software/webx/engine/stencils/infrastructure/jsons/instances/layers/instructions/assignments/assignables/executions/executes"
	json_inits "github.com/steve-care-software/webx/engine/stencils/infrastructure/jsons/instances/layers/instructions/assignments/assignables/executions/inits"
	json_retrieves "github.com/steve-care-software/webx/engine/stencils/infrastructure/jsons/instances/layers/instructions/assignments/assignables/executions/retrieves"
)

// Adapter represents an adapter
type Adapter struct {
	executeAdapter  *json_executes.Adapter
	initAdapter     *json_inits.Adapter
	retrieveAdapter *json_retrieves.Adapter
	builder         executions.Builder
	contentBuilder  executions.ContentBuilder
}

func createAdapter(
	executeAdapter *json_executes.Adapter,
	initAdapter *json_inits.Adapter,
	retrieveAdapter *json_retrieves.Adapter,
	builder executions.Builder,
	contentBuilder executions.ContentBuilder,
) executions.Adapter {
	out := Adapter{
		executeAdapter:  executeAdapter,
		initAdapter:     initAdapter,
		retrieveAdapter: retrieveAdapter,
		builder:         builder,
		contentBuilder:  contentBuilder,
	}

	return &out
}

// ToBytes converts an instance to bytes
func (app *Adapter) ToBytes(ins executions.Execution) ([]byte, error) {
	str := app.ExecutionToStruct(ins)
	js, err := json.Marshal(str)
	if err != nil {
		return nil, err
	}

	return js, nil
}

// ToInstance converts a bytes to instance
func (app *Adapter) ToInstance(data []byte) (executions.Execution, error) {
	ins := new(Execution)
	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err := decoder.Decode(ins)
	if err != nil {
		return nil, err
	}

	return app.StructToExecution(*ins)
}

// ExecutionToStruct converts an execution to struct
func (app *Adapter) ExecutionToStruct(ins executions.Execution) Execution {
	str := app.ContentToStruct(ins.Content())
	return Execution{
		Executable: ins.Executable(),
		Content:    str,
	}
}

// StructToExecution converts a struct to execution
func (app *Adapter) StructToExecution(str Execution) (executions.Execution, error) {
	content, err := app.StructToContent(str.Content)
	if err != nil {
		return nil, err
	}

	return app.builder.Create().
		WithExecutable(str.Executable).
		WithContent(content).
		Now()
}

// StructToContent converts a struct to content
func (app *Adapter) StructToContent(str Content) (executions.Content, error) {
	builder := app.contentBuilder.Create()
	if str.Amount != "" {
		builder.WithAmount(str.Amount)
	}

	if str.Begin != "" {
		builder.WithBegin(str.Begin)
	}

	if str.Execute != nil {
		execute, err := app.executeAdapter.StructToExecute(*str.Execute)
		if err != nil {
			return nil, err
		}

		builder.WithExecute(execute)
	}

	if str.Head != "" {
		builder.WithHead(str.Head)
	}

	if str.Init != nil {
		init, err := app.initAdapter.StructToInit(*str.Init)
		if err != nil {
			return nil, err
		}

		builder.WithInit(init)
	}

	if str.Retrieve != nil {
		retrieve, err := app.retrieveAdapter.StructToRetrieve(*str.Retrieve)
		if err != nil {
			return nil, err
		}

		builder.WithRetrieve(retrieve)
	}

	if str.IsList {
		builder.IsList()
	}

	return builder.Now()
}

// ContentToStruct converts a content to struct
func (app *Adapter) ContentToStruct(ins executions.Content) Content {
	out := Content{}
	if ins.IsAmount() {
		out.Amount = ins.Amount()
	}

	if ins.IsBegin() {
		out.Begin = ins.Begin()
	}

	if ins.IsExecute() {
		execute := app.executeAdapter.ExecuteToStruct(ins.Execute())
		out.Execute = &execute
	}

	if ins.IsHead() {
		out.Head = ins.Head()
	}

	if ins.IsInit() {
		init := app.initAdapter.InitToStruct(ins.Init())
		out.Init = &init
	}

	if ins.IsRetrieve() {
		retrieve := app.retrieveAdapter.RetrieveToStruct(ins.Retrieve())
		out.Retrieve = &retrieve
	}

	if ins.IsList() {
		out.IsList = true
	}
	return out
}
