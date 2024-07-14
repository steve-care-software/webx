package executions

import (
	json_merges "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/layers/instructions/executions/merges"
)

// Execution represents an execution
type Execution struct {
	Commit   string             `json:"commit"`
	Rollback string             `json:"rollback"`
	Cancel   string             `json:"cancel"`
	Merge    *json_merges.Merge `json:"merge"`
}
