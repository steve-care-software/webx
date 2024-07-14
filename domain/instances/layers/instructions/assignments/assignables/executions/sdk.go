package executions

import (
	"github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables/executions/amounts"
	"github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables/executions/begins"
	"github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables/executions/executes"
	"github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables/executions/heads"
	"github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables/executions/inits"
	"github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables/executions/retrievealls"
	"github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables/executions/retrieveats"
	"github.com/steve-care-software/historydb/domain/hash"
)

// Builder represents an execution builder
type Builder interface {
	Create() Builder
	IsList() Builder
	Now() (Execution, error)
}

// Execution represents an execution instruction
type Execution interface {
	Hash() hash.Hash
	IsList() bool
	List() string
	IsInit() bool
	Init() inits.Init
	IsBegin() bool
	Begin() begins.Begin
	IsExecute() bool
	Execute() executes.Execute
	IsRetrieveAll() bool
	RetrieveAll() retrievealls.RetrieveAll
	IsRetrieveAt() bool
	RetrieveAt() retrieveats.RetrieveAt
	IsAmount() bool
	Amount() amounts.Amount
	IsHead() bool
	Head() heads.Head
}
