package lists

import (
	json_deletes "github.com/steve-care-software/datastencil/stencils/infrastructure/jsons/instances/layers/instructions/lists/deletes"
	json_inserts "github.com/steve-care-software/datastencil/stencils/infrastructure/jsons/instances/layers/instructions/lists/inserts"
)

// List represents a list
type List struct {
	Delete *json_deletes.Delete `json:"delete"`
	Insert *json_inserts.Insert `json:"insert"`
}
