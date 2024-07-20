package lists

import (
	json_fetches "github.com/steve-care-software/webx/engine/stencils/infrastructure/jsons/instances/layers/instructions/assignments/assignables/lists/fetches"
)

// List represents a list
type List struct {
	Fetch  *json_fetches.Fetch `json:"fetch"`
	Length string              `json:"length"`
	Create string              `json:"create"`
}
