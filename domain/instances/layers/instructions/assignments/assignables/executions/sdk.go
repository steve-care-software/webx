package executions

import (
	"github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables/executions/amounts"
	"github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables/executions/begins"
	"github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables/executions/executes"
	"github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables/executions/heads"
	"github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables/executions/inits"
	"github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables/executions/retrieves"
	"github.com/steve-care-software/historydb/domain/hash"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// Adapter represents the execution adapter
type Adapter interface {
	ToBytes(ins Execution) ([]byte, error)
	ToInstance(bytes []byte) (Execution, error)
}

// Builder represents an execution builder
type Builder interface {
	Create() Builder
	WithList(list string) Builder
	WithInit(init inits.Init) Builder
	WithBegin(begin begins.Begin) Builder
	WithExecute(execute executes.Execute) Builder
	WithRetrieve(retrieve retrieves.Retrieve) Builder
	WithAmount(amount amounts.Amount) Builder
	WithHead(head heads.Head) Builder
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
	IsRetrieve() bool
	Retrieve() retrieves.Retrieve
	IsAmount() bool
	Amount() amounts.Amount
	IsHead() bool
	Head() heads.Head
}
