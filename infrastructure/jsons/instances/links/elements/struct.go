package elements

import (
	json_conditions "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/links/elements/conditions"
)

// Element represents an element
type Element struct {
	Layer     []string                   `json:"layer"`
	Condition *json_conditions.Condition `json:"condition"`
}
