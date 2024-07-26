package jsons

import (
	"encoding/base64"
	"encoding/json"

	"github.com/steve-care-software/webx/engine/databases/entities/domain/hash"
	"github.com/steve-care-software/webx/engine/states/domain/databases/commits"
	"github.com/steve-care-software/webx/engine/states/domain/databases/commits/executions"
	"github.com/steve-care-software/webx/engine/states/domain/databases/commits/executions/chunks"
)

type commitAdapter struct {
	commitBuilder     commits.Builder
	executionsBuilder executions.Builder
	executionBuilder  executions.ExecutionBuilder
	chunkBuilder      chunks.Builder
	hashAdapter       hash.Adapter
}

func createCommitAdapter(
	commitBuilder commits.Builder,
	executionsBuilder executions.Builder,
	executionBuilder executions.ExecutionBuilder,
	chunkBuilder chunks.Builder,
	hashAdapter hash.Adapter,
) commits.Adapter {
	out := commitAdapter{
		commitBuilder:     commitBuilder,
		executionsBuilder: executionsBuilder,
		executionBuilder:  executionBuilder,
		chunkBuilder:      chunkBuilder,
		hashAdapter:       hashAdapter,
	}

	return &out
}

// ToBytes converts commit to bytes
func (app *commitAdapter) ToBytes(ins commits.Commit) ([]byte, error) {
	executionsStrList := []execution{}
	executionsList := ins.Executions().List()
	for _, oneExecution := range executionsList {
		executionStr := execution{}
		if oneExecution.IsBytes() {
			executionStr.Data = base64.StdEncoding.EncodeToString(oneExecution.Bytes())
		}

		if oneExecution.IsChunk() {
			chunkIns := oneExecution.Chunk()
			executionStr.Chunk = &chunk{
				Path:        chunkIns.Path(),
				Fingerprint: chunkIns.FingerPrint().String(),
			}
		}

		executionsStrList = append(executionsStrList, executionStr)
	}

	commitStr := commit{
		Executions: executionsStrList,
	}

	if ins.HasParent() {
		commitStr.Parent = ins.Parent().String()
	}

	return json.Marshal(commitStr)
}

// ToInstance converts bytes to instance
func (app *commitAdapter) ToInstance(bytes []byte) (commits.Commit, error) {
	ptr := new(commit)
	err := json.Unmarshal(bytes, ptr)
	if err != nil {
		return nil, err
	}

	executionsList := []executions.Execution{}
	for _, oneExecutionStr := range ptr.Executions {
		executionBuilder := app.executionBuilder.Create()
		if oneExecutionStr.Data != "" {
			decodedBytes, err := base64.StdEncoding.DecodeString(oneExecutionStr.Data)
			if err != nil {
				return nil, err
			}

			executionBuilder.WithBytes(decodedBytes)
		}

		if oneExecutionStr.Chunk != nil {
			chkStr := oneExecutionStr.Chunk
			pFingerprint, err := app.hashAdapter.FromString(chkStr.Fingerprint)
			if err != nil {
				return nil, err
			}

			chunk, err := app.chunkBuilder.Create().
				WithPath(chkStr.Path).
				WithFingerPrint(*pFingerprint).
				Now()

			if err != nil {
				return nil, err
			}

			executionBuilder.WithChunk(chunk)
		}

		execution, err := executionBuilder.Now()
		if err != nil {
			return nil, err
		}

		executionsList = append(executionsList, execution)
	}

	excecutionsIns, err := app.executionsBuilder.Create().WithList(executionsList).Now()
	if err != nil {
		return nil, err
	}

	builder := app.commitBuilder.Create().WithExecutions(excecutionsIns)
	if ptr.Parent != "" {
		pParent, err := app.hashAdapter.FromString(ptr.Parent)
		if err != nil {
			return nil, err
		}

		builder.WithParent(*pParent)
	}

	return builder.Now()
}
