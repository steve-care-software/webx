package elements

import (
	json_conditions "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/links/elements/conditions"
	json_layers "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/links/elements/layers"
)

// Element represents an element
type Element struct {
	Layer     json_layers.Layer          `json:"layer"`
	Condition *json_conditions.Condition `json:"condition"`
}
