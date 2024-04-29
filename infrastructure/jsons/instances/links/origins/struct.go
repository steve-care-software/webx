package origins

import (
	json_operators "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/links/origins/operators"
	json_resources "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/links/origins/resources"
)

// Origin represents an origin
type Origin struct {
	Resource json_resources.Resource `json:"resource"`
	Operator json_operators.Operator `json:"operator"`
	Next     Value                   `json:"value"`
}

// Value represents a value
type Value struct {
	Resource *json_resources.Resource `json:"resource"`
	Origin   *Origin                  `json:"origin"`
}
