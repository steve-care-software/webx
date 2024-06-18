package pointers

import (
	json_conditions "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/pointers/conditions"
)

// Pointer represents a pointer
type Pointer struct {
	Path      []string                   `json:"path"`
	IsActive  bool                       `json:"is_active"`
	Loader    *json_conditions.Condition `json:"loader"`
	Canceller *json_conditions.Condition `json:"canceller"`
}
