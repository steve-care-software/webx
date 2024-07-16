package executions

import (
	"bytes"
	"encoding/json"

	"github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables/executions"
	json_amounts "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/layers/instructions/assignments/assignables/executions/amounts"
	json_begins "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/layers/instructions/assignments/assignables/executions/begins"
	json_executes "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/layers/instructions/assignments/assignables/executions/executes"
	json_heads "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/layers/instructions/assignments/assignables/executions/heads"
	json_inits "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/layers/instructions/assignments/assignables/executions/inits"
	json_retrieves "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/layers/instructions/assignments/assignables/executions/retrieves"
)

// Adapter represents an adapter
type Adapter struct {
	amountAdapter   *json_amounts.Adapter
	beginAdapter    *json_begins.Adapter
	executeAdapter  *json_executes.Adapter
	headAdapter     *json_heads.Adapter
	initAdapter     *json_inits.Adapter
	retrieveAdapter *json_retrieves.Adapter
	builder         executions.Builder
}

func createAdapter(
	amountAdapter *json_amounts.Adapter,
	beginAdapter *json_begins.Adapter,
	executeAdapter *json_executes.Adapter,
	headAdapter *json_heads.Adapter,
	initAdapter *json_inits.Adapter,
	retrieveAdapter *json_retrieves.Adapter,
	builder executions.Builder,
) executions.Adapter {
	out := Adapter{
		amountAdapter:   amountAdapter,
		beginAdapter:    beginAdapter,
		executeAdapter:  executeAdapter,
		headAdapter:     headAdapter,
		initAdapter:     initAdapter,
		retrieveAdapter: retrieveAdapter,
		builder:         builder,
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
	out := Execution{}
	if ins.IsAmount() {
		amount := app.amountAdapter.AmountToStruct(ins.Amount())
		out.Amount = &amount
	}

	if ins.IsBegin() {
		begin := app.beginAdapter.BeginToStruct(ins.Begin())
		out.Begin = &begin
	}

	if ins.IsExecute() {
		execute := app.executeAdapter.ExecuteToStruct(ins.Execute())
		out.Execute = &execute
	}

	if ins.IsHead() {
		head := app.headAdapter.HeadToStruct(ins.Head())
		out.Head = &head
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
		out.List = ins.List()
	}
	return out
}

// StructToExecution converts a struct to execution
func (app *Adapter) StructToExecution(str Execution) (executions.Execution, error) {
	builder := app.builder.Create()
	if str.Amount != nil {
		amount, err := app.amountAdapter.StructToAmount(*str.Amount)
		if err != nil {
			return nil, err
		}

		builder.WithAmount(amount)
	}

	if str.Begin != nil {
		begin, err := app.beginAdapter.StructToBegin(*str.Begin)
		if err != nil {
			return nil, err
		}

		builder.WithBegin(begin)
	}

	if str.Execute != nil {
		execute, err := app.executeAdapter.StructToExecute(*str.Execute)
		if err != nil {
			return nil, err
		}

		builder.WithExecute(execute)
	}

	if str.Head != nil {
		head, err := app.headAdapter.StructToHead(*str.Head)
		if err != nil {
			return nil, err
		}

		builder.WithHead(head)
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

	if str.List != "" {
		builder.WithList(str.List)
	}

	return builder.Now()
}
