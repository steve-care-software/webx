package executions

import (
	json_results "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/executions/results"
	json_layers "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/layers"
)

// Execution represents an execution
type Execution struct {
	Input  []byte               `json:"input"`
	Source *json_layers.Layer   `json:"source"`
	Result *json_results.Result `json:"result"`
}
