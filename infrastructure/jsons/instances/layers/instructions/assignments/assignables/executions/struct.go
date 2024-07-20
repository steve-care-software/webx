package executions

import (
	json_amounts "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/layers/instructions/assignments/assignables/executions/amounts"
	json_begins "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/layers/instructions/assignments/assignables/executions/begins"
	json_executes "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/layers/instructions/assignments/assignables/executions/executes"
	json_heads "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/layers/instructions/assignments/assignables/executions/heads"
	json_inits "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/layers/instructions/assignments/assignables/executions/inits"
	json_retrieves "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/layers/instructions/assignments/assignables/executions/retrieves"
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
	Begin    *json_begins.Begin       `json:"begin"`
	Execute  *json_executes.Execute   `json:"execute"`
	Retrieve *json_retrieves.Retrieve `json:"retrieve"`
	Amount   *json_amounts.Amount     `json:"amount"`
	Head     *json_heads.Head         `json:"head"`
}
