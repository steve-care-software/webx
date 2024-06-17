package success

import (
	json_outputs "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/commands/results/success/outputs"
	json_kinds "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/pointers/resources/logics/bridges/layers/outputs/kinds"
)

// Success represents a success
type Success struct {
	Output json_outputs.Output `json:"output"`
	Kind   json_kinds.Kind     `json:"kind"`
}
