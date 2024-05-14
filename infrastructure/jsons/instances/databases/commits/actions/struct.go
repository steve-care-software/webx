package actions

import (
	json_modifications "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/databases/commits/actions/modifications"
)

// Action represents an action
type Action struct {
	Path          []string                          `json:"path"`
	Modifications []json_modifications.Modification `json:"modifications"`
}
