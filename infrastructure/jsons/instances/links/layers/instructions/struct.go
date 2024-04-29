package instructions

import (
	json_accounts "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/links/layers/instructions/accounts"
	json_assignments "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/links/layers/instructions/assignments"
	json_databases "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/links/layers/instructions/databases"
)

// Instruction represents an instruction
type Instruction struct {
	Stop       bool                         `json:"stop"`
	RaiseError *uint                        `json:"raise_error"`
	Condition  *Condition                   `json:"condition"`
	Assignment *json_assignments.Assignment `json:"assignment"`
	Account    *json_accounts.Account       `json:"account"`
	Database   *json_databases.Database     `json:"database"`
}

// Condition represents a condition
type Condition struct {
	Variable     string        `json:"variable"`
	Instructions []Instruction `json:"instructions"`
}
