package executions

import (
	"bytes"
	"encoding/json"

	"github.com/steve-care-software/webx/engine/stencils/domain/instances/layers/instructions/executions"
	json_merges "github.com/steve-care-software/webx/engine/stencils/infrastructure/jsons/instances/layers/instructions/executions/merges"
)

// Adapter represents an execution adapter
type Adapter struct {
	mergeAdapter   *json_merges.Adapter
	builder        executions.Builder
	contentBuilder executions.ContentBuilder
}

func createAdapter(
	mergeAdapter *json_merges.Adapter,
	builder executions.Builder,
	contentBuilder executions.ContentBuilder,
) executions.Adapter {
	out := Adapter{
		mergeAdapter:   mergeAdapter,
		builder:        builder,
		contentBuilder: contentBuilder,
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
	content := app.ContentToStruct(ins.Content())
	return Execution{
		Executable: ins.Executable(),
		Content:    content,
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

// ContentToStruct converts a content to struct
func (app *Adapter) ContentToStruct(ins executions.Content) Content {
	out := Content{}
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

// StructToContent converts a struct to content
func (app *Adapter) StructToContent(str Content) (executions.Content, error) {
	builder := app.contentBuilder.Create()
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
