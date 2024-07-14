package executions

import (
	"encoding/json"

	"github.com/steve-care-software/datastencil/domain/instances/layers/instructions/executions"
	json_merges "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/layers/instructions/executions/merges"
)

// Adapter represents an execution adapter
type Adapter struct {
	mergeAdapter *json_merges.Adapter
	builder      executions.Builder
}

func createAdapter(
	mergeAdapter *json_merges.Adapter,
	builder executions.Builder,
) executions.Adapter {
	out := Adapter{
		mergeAdapter: mergeAdapter,
		builder:      builder,
	}

	return &out
}

// ToBytes converts instance to bytes
func (app *Adapter) ToBytes(ins executions.Execution) ([]byte, error) {
	str := app.ExecutionToStruct(ins)
	js, err := json.Marshal(str)
	if err != nil {
		return nil, err
	}

	return js, nil
}

// ToInstance converts bytes to instance
func (app *Adapter) ToInstance(bytes []byte) (executions.Execution, error) {
	ins := new(Execution)
	err := json.Unmarshal(bytes, ins)
	if err != nil {
		return nil, err
	}

	return app.StructToExecution(*ins)
}

// ExecutionToStruct converts an execution to struct
func (app *Adapter) ExecutionToStruct(ins executions.Execution) Execution {
	out := Execution{}
	if ins.IsCommit() {
		out.Commit = ins.Commit()
	}

	if ins.IsRollback() {
		out.Rollback = ins.Rollback()
	}

	if ins.IsCancel() {
		out.Cancel = ins.Cancel()
	}

	if ins.IsMerge() {
		merge := app.mergeAdapter.MergeToStruct(ins.Merge())
		out.Merge = &merge
	}

	return out
}

// StructToExecution converts a struct to execution
func (app *Adapter) StructToExecution(str Execution) (executions.Execution, error) {
	builder := app.builder.Create()
	if str.Commit != "" {
		builder.WithCommit(str.Commit)
	}

	if str.Rollback != "" {
		builder.WithRollback(str.Rollback)
	}

	if str.Cancel != "" {
		builder.WithCancel(str.Cancel)
	}

	if str.Merge != nil {
		merge, err := app.mergeAdapter.StructToMerge(*str.Merge)
		if err != nil {
			return nil, err
		}

		builder.WithMerge(merge)
	}

	return builder.Now()
}
