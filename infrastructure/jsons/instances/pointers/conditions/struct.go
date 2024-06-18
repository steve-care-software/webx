package conditions

import (
	json_operators "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/pointers/conditions/operators"
)

// Condition represents a condition
type Condition struct {
	Resource    Resource     `json:"resource"`
	Comparisons []Comparison `json:"comparisons"`
}

// Resource represents a resource
type Resource struct {
	Path         []string `json:"path"`
	MustBeLoaded bool     `json:"must_be_loaded"`
}

// Comparison represents a comparison
type Comparison struct {
	Operator  json_operators.Operator `json:"operator"`
	Condition Condition               `json:"condition"`
}
