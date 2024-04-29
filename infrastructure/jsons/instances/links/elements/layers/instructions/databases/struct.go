package databases

import (
	json_deletes "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/links/elements/layers/instructions/databases/deletes"
	json_inserts "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/links/elements/layers/instructions/databases/inserts"
	json_reverts "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/links/elements/layers/instructions/databases/reverts"
)

// Database represents a database
type Database struct {
	Insert *json_inserts.Insert `json:"insert"`
	Delete *json_deletes.Delete `json:"delete"`
	Commit string               `json:"commit"`
	Cancel string               `json:"cancel"`
	Revert *json_reverts.Revert `json:"revert"`
}
