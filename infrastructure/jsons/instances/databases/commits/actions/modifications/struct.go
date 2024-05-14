package modifications

import (
	json_deletes "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/databases/commits/actions/modifications/deletes"
)

// Modification creates a new modification for tests
type Modification struct {
	Insert string               `json:"insert"`
	Delete *json_deletes.Delete `json:"delete"`
}
