package memories

import (
	"errors"
	"fmt"
	"path/filepath"

	"github.com/steve-care-software/webx/engine/states/domain/hash"
	"github.com/steve-care-software/webx/engine/stencils/domain/instances/executions"
)

type executionRepository struct {
	memory  map[string]map[string]executions.Execution
	builder executions.Builder
}

func createExecutionRepository(
	memory map[string]map[string]executions.Execution,
	builder executions.Builder,
) executions.Repository {
	out := executionRepository{
		memory:  memory,
		builder: builder,
	}

	return &out
}

// RetrieveAll retrieves executions
func (app *executionRepository) RetrieveAll(dbPath []string, hashes []hash.Hash) (executions.Executions, error) {
	keyname := filepath.Join(dbPath...)
	if execMap, ok := app.memory[keyname]; ok {
		list := []executions.Execution{}
		for _, oneHash := range hashes {
			if execIns, ok := execMap[oneHash.String()]; ok {
				list = append(list, execIns)
				continue
			}

			str := fmt.Sprintf("the database path (%s) contains no execution for the provided hash (%s)", keyname, oneHash.String())
			return nil, errors.New(str)
		}

		return app.builder.Create().WithList(list).Now()
	}

	str := fmt.Sprintf("the database path (%s) does not exists", keyname)
	return nil, errors.New(str)
}
