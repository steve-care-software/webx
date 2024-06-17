package instructions

import (
	json_assignments "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/pointers/resources/logics/bridges/layers/instructions/assignments"
	json_databases "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/pointers/resources/logics/bridges/layers/instructions/databases"
	json_lists "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/pointers/resources/logics/bridges/layers/instructions/lists"
)

// Instruction represents an instruction
type Instruction struct {
	Stop       bool                         `json:"stop"`
	RaiseError *uint                        `json:"raise_error"`
	Condition  *Condition                   `json:"condition"`
	Loop       *Loop                        `json:"loop"`
	Assignment *json_assignments.Assignment `json:"assignment"`
	Database   *json_databases.Database     `json:"database"`
	List       *json_lists.List             `json:"list"`
}

// Condition represents a condition
type Condition struct {
	Variable     string        `json:"variable"`
	Instructions []Instruction `json:"instructions"`
}

// Loop represents a loop
type Loop struct {
	Amount       string        `json:"amount"`
	Instructions []Instruction `json:"instructions"`
}
