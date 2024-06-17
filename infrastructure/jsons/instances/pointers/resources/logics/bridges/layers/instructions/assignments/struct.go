package assignments

import (
	json_assignables "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/pointers/resources/logics/bridges/layers/instructions/assignments/assignables"
)

// Assignment represents an assignment
type Assignment struct {
	Name       string                      `json:"name"`
	Assignable json_assignables.Assignable `json:"assignable"`
}
