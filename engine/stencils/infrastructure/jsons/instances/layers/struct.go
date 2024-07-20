package layers

import (
	json_instructions "github.com/steve-care-software/webx/engine/stencils/infrastructure/jsons/instances/layers/instructions"
	json_output "github.com/steve-care-software/webx/engine/stencils/infrastructure/jsons/instances/layers/outputs"
	json_references "github.com/steve-care-software/webx/engine/stencils/infrastructure/jsons/instances/layers/references"
)

// Layer represents the layer
type Layer struct {
	Instructions []json_instructions.Instruction `json:"instructions"`
	Output       json_output.Output              `json:"output"`
	Input        string                          `json:"input"`
	References   []json_references.Reference     `json:"references"`
}
