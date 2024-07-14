package executes

import json_inputs "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/layers/instructions/assignments/assignables/executions/executes/inputs"

// Execute represents an execute
type Execute struct {
	Context string             `json:"context"`
	Input   json_inputs.Input  `json:"input"`
	Return  string             `json:"return"`
	Layer   *json_inputs.Input `json:"layer"`
}
