package executions

import (
	"bytes"
	"encoding/json"

	"github.com/steve-care-software/webx/engine/stencils/domain/instances/executions"
	json_results "github.com/steve-care-software/webx/engine/stencils/infrastructure/jsons/instances/executions/results"
	json_layers "github.com/steve-care-software/webx/engine/stencils/infrastructure/jsons/instances/layers"
)

// Adapter represents an adapter
type Adapter struct {
	resultAdapter     *json_results.Adapter
	layerAdapter      *json_layers.Adapter
	builder           executions.Builder
	executionsBuilder executions.ExecutionBuilder
}

func createAdapter(
	resultAdapter *json_results.Adapter,
	layerAdapter *json_layers.Adapter,
	builder executions.Builder,
	executionsBuilder executions.ExecutionBuilder,
) executions.Adapter {
	out := Adapter{
		resultAdapter:     resultAdapter,
		layerAdapter:      layerAdapter,
		builder:           builder,
		executionsBuilder: executionsBuilder,
	}

	return &out
}

// InstanceToBytes converts instance to bytes
func (app *Adapter) InstanceToBytes(ins executions.Execution) ([]byte, error) {
	ptr, err := app.ExecutionToStruct(ins)
	if err != nil {
		return nil, err
	}

	js, err := json.Marshal(ptr)
	if err != nil {
		return nil, err
	}

	return js, nil
}

// BytesToInstance converts bytes to instance
func (app *Adapter) BytesToInstance(data []byte) (executions.Execution, error) {
	ins := new(Execution)
	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err := decoder.Decode(ins)
	if err != nil {
		return nil, err
	}

	return app.StructToExecution(*ins)
}

// InstancesToBytes converts instances to bytes
func (app *Adapter) InstancesToBytes(ins executions.Executions) ([]byte, error) {
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

// BytesToInstances converts bytes to instances
func (app *Adapter) BytesToInstances(bytes []byte) (executions.Executions, error) {
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
	for _, oneStr := range str {
		ins, err := app.StructToExecution(oneStr)
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
	source, err := app.layerAdapter.LayerToStruct(ins.Source())
	if err != nil {
		return nil, err
	}

	result, err := app.resultAdapter.ResultToStruct(ins.Result())
	if err != nil {
		return nil, err
	}

	return &Execution{
		Input:  ins.Input(),
		Source: *source,
		Result: *result,
	}, nil
}

// StructToExecution converts a struct to execution
func (app *Adapter) StructToExecution(str Execution) (executions.Execution, error) {
	source, err := app.layerAdapter.StructToLayer(str.Source)
	if err != nil {
		return nil, err
	}

	result, err := app.resultAdapter.StructToResult(str.Result)
	if err != nil {
		return nil, err
	}

	return app.executionsBuilder.Create().
		WithInput(str.Input).
		WithSource(source).
		WithResult(result).
		Now()
}
