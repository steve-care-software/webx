package bridges

import (
	json_layers "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/pointers/resources/logics/bridges/layers"
)

// Bridge represents a bridge
type Bridge struct {
	Path  []string          `json:"path"`
	Layer json_layers.Layer `json:"layer"`
}
