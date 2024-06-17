package lists

import (
	json_fetches "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/pointers/resources/logics/bridges/layers/instructions/assignments/assignables/lists/fetches"
)

// List represents a list
type List struct {
	Fetch  *json_fetches.Fetch `json:"fetch"`
	Length string              `json:"length"`
	Create string              `json:"create"`
}
