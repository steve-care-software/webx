package outputs

import (
	json_kinds "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/links/elements/layers/outputs/kinds"
)

// Output represents an output
type Output struct {
	Variable string          `json:"variable"`
	Kind     json_kinds.Kind `json:"kind"`
	Execute  []string        `json:"execute"`
}
