package jsons

import (
	json_executions "github.com/steve-care-software/webx/engine/vms/infrastructure/jsons/instances/executions"
)

// Session represents a session
type Session struct {
	Executions [][]json_executions.Execution `json:"executions"`
}
