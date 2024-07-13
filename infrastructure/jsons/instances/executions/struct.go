package executions

import (
	json_layers "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/executions/layers"
	json_databases "github.com/steve-care-software/historydb/infrastructure/jsons"
)

// Execution represents an execution
type Execution struct {
	Layer    json_layers.Layer       `json:"layer"`
	Database json_databases.Database `json:"database"`
}
