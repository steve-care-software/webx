package layers

import (
	json_results "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/executions/links/layers/results"
	json_layers "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/layers"
)

// Layer represents a layer
type Layer struct {
	Input  string              `json:"input"`
	Source json_layers.Layer   `json:"source"`
	Result json_results.Result `json:"result"`
}
