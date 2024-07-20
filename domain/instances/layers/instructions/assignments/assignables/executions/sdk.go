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

// NewBuilder creates a new builder for tests
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// NewContentBuilder creates a new content builder instance
func NewContentBuilder() ContentBuilder {
	hashAdapter := hash.NewAdapter()
	return createContentBuilder(
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
	WithExecutable(executable string) Builder
	WithContent(content Content) Builder
	Now() (Execution, error)
}

// Execution represents an execution instruction
type Execution interface {
	Hash() hash.Hash
	Executable() string
	Content() Content
}

// ContentBuilder represents a content builder
type ContentBuilder interface {
	Create() ContentBuilder
	WithInit(init inits.Init) ContentBuilder
	WithBegin(begin begins.Begin) ContentBuilder
	WithExecute(execute executes.Execute) ContentBuilder
	WithRetrieve(retrieve retrieves.Retrieve) ContentBuilder
	WithAmount(amount amounts.Amount) ContentBuilder
	WithHead(head heads.Head) ContentBuilder
	IsList() ContentBuilder
	Now() (Content, error)
}

// Content represents the content execution
type Content interface {
	Hash() hash.Hash
	IsList() bool
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
