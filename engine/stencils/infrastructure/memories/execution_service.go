package memories

import (
	"path/filepath"

	"github.com/steve-care-software/webx/engine/stencils/domain/instances/executions"
)

type executionService struct {
	memory map[string]map[string]executions.Execution
}

func createExecutionService(
	memory map[string]map[string]executions.Execution,
) executions.Service {
	out := executionService{
		memory: memory,
	}

	return &out
}

// Save saves an execution
func (app *executionService) Save(dbPath []string, ins executions.Execution) error {
	keyname := filepath.Join(dbPath...)
	if _, ok := app.memory[keyname]; !ok {
		app.memory[keyname] = map[string]executions.Execution{}
	}

	hashStr := ins.Hash().String()
	app.memory[keyname][hashStr] = ins
	return nil
}
