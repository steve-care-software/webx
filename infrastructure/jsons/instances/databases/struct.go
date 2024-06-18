package databases

import (
	json_commits "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/databases/commits"
	json_heads "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/databases/heads"
)

// Database represents a database
type Database struct {
	Head   json_heads.Head     `json:"head"`
	Commit json_commits.Commit `json:"commit"`
}
