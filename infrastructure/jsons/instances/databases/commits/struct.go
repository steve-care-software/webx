package commits

import (
	json_actions "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/databases/commits/actions"
)

// Commit represents the commit
type Commit struct {
	Description string                `json:"description"`
	Actions     []json_actions.Action `json:"actions"`
	Parent      string                `json:"parent"`
}
