package instructions

import (
	json_assignments "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/links/elements/layers/instructions/assignments"
)

// Instruction represents an instruction
type Instruction struct {
	Stop       bool                         `json:"stop"`
	RaiseError *uint                        `json:"raise_error"`
	Condition  *Condition                   `json:"condition"`
	Assignment *json_assignments.Assignment `json:"assignment"`
}

// Condition represents a condition
type Condition struct {
	Variable     string        `json:"variable"`
	Instructions []Instruction `json:"instructions"`
}
