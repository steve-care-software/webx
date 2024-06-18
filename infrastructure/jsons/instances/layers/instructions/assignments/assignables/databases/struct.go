package databases

import (
	json_actions "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/layers/instructions/assignments/assignables/databases/actions"
	json_commits "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/layers/instructions/assignments/assignables/databases/commits"
	json_databases "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/layers/instructions/assignments/assignables/databases/databases"
	json_deletes "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/layers/instructions/assignments/assignables/databases/deletes"
	json_modifications "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/layers/instructions/assignments/assignables/databases/modifications"
	json_retrieves "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/layers/instructions/assignments/assignables/databases/retrieves"
)

// Database represents a database
type Database struct {
	Database     *json_databases.Database         `json:"database"`
	Commit       *json_commits.Commit             `json:"commit"`
	Action       *json_actions.Action             `json:"action"`
	Modification *json_modifications.Modification `json:"modification"`
	Delete       *json_deletes.Delete             `json:"delete"`
	Retrieve     *json_retrieves.Retrieve         `json:"retrieve"`
}
