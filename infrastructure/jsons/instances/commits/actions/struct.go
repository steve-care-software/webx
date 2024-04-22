package actions

import (
	"github.com/steve-care-software/datastencil/infrastructure/jsons/instances/commits/actions/pointers"
	"github.com/steve-care-software/datastencil/infrastructure/jsons/instances/commits/actions/resources"
)

// Action represents an action
type Action struct {
	Insert *resources.Resource `json:"insert"`
	Delete *pointers.Pointer   `json:"delete"`
}
