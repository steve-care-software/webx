package executions

import (
	json_executes "github.com/steve-care-software/webx/engine/vms/infrastructure/jsons/instances/layers/instructions/assignments/assignables/executions/executes"
	json_inits "github.com/steve-care-software/webx/engine/vms/infrastructure/jsons/instances/layers/instructions/assignments/assignables/executions/inits"
	json_retrieves "github.com/steve-care-software/webx/engine/vms/infrastructure/jsons/instances/layers/instructions/assignments/assignables/executions/retrieves"
)

// Execution represents an execution
type Execution struct {
	Executable string  `json:"executable"`
	Content    Content `json:"content"`
}

// Content represents content
type Content struct {
	IsList   bool                     `json:"is_list"`
	Init     *json_inits.Init         `json:"init"`
	Begin    string                   `json:"begin"`
	Execute  *json_executes.Execute   `json:"execute"`
	Retrieve *json_retrieves.Retrieve `json:"retrieve"`
	Amount   string                   `json:"amount"`
	Head     string                   `json:"head"`
}
