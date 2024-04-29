package conditions

import (
	json_resources "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/links/elements/conditions/resources"
)

// Condition represents a condition
type Condition struct {
	Resource json_resources.Resource `json:"resource"`
	Next     *Value                  `json:"next"`
}

// Value represents a condition value
type Value struct {
	Resource  *json_resources.Resource `json:"resource"`
	Condition *Condition               `json:"condition"`
}
