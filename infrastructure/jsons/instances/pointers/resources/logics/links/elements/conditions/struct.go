package conditions

import (
	json_resources "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/pointers/resources/logics/links/elements/conditions/resources"
)

// Condition represents a condition
type Condition struct {
	Resource json_resources.Resource `json:"resource"`
	Next     *Condition              `json:"next"`
}
