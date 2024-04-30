package operators

import (
	json_integers "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/queries/conditions/operators/integers"
	json_relationals "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/queries/conditions/operators/relationals"
)

// Operator represents an operator
type Operator struct {
	Integer    *json_integers.Integer       `json:"integer"`
	Relational *json_relationals.Relational `json:"relational"`
	IsEqual    bool                         `json:"is_equal"`
}
