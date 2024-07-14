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

type execution struct {
	hash     hash.Hash
	list     string
	init     inits.Init
	begin    begins.Begin
	execute  executes.Execute
	retrieve retrieves.Retrieve
	amount   amounts.Amount
	head     heads.Head
}

func createExecutionWithList(
	hash hash.Hash,
	list string,
) Execution {
	return createExecutionInternally(hash, list, nil, nil, nil, nil, nil, nil)
}

func createExecutionWithInit(
	hash hash.Hash,
	init inits.Init,
) Execution {
	return createExecutionInternally(hash, "", init, nil, nil, nil, nil, nil)
}

func createExecutionWithBegin(
	hash hash.Hash,
	begin begins.Begin,
) Execution {
	return createExecutionInternally(hash, "", nil, begin, nil, nil, nil, nil)
}

func createExecutionWithExecute(
	hash hash.Hash,
	execute executes.Execute,
) Execution {
	return createExecutionInternally(hash, "", nil, nil, execute, nil, nil, nil)
}

func createExecutionWithRetrieve(
	hash hash.Hash,
	retrieve retrieves.Retrieve,
) Execution {
	return createExecutionInternally(hash, "", nil, nil, nil, retrieve, nil, nil)
}

func createExecutionWithAmount(
	hash hash.Hash,
	amount amounts.Amount,
) Execution {
	return createExecutionInternally(hash, "", nil, nil, nil, nil, amount, nil)
}

func createExecutionWithHead(
	hash hash.Hash,
	head heads.Head,
) Execution {
	return createExecutionInternally(hash, "", nil, nil, nil, nil, nil, head)
}

func createExecutionInternally(
	hash hash.Hash,
	list string,
	init inits.Init,
	begin begins.Begin,
	execute executes.Execute,
	retrieve retrieves.Retrieve,
	amount amounts.Amount,
	head heads.Head,
) Execution {
	out := execution{
		hash:     hash,
		list:     list,
		init:     init,
		begin:    begin,
		execute:  execute,
		retrieve: retrieve,
		amount:   amount,
		head:     head,
	}

	return &out
}

// Hash returns the hash
func (obj *execution) Hash() hash.Hash {
	return obj.hash
}

// IsList returns true if there is a list, false otherwise
func (obj *execution) IsList() bool {
	return obj.list != ""
}

// List returns the list, if any
func (obj *execution) List() string {
	return obj.list
}

// IsInit returns true if there is an init, false otherwise
func (obj *execution) IsInit() bool {
	return obj.init != nil
}

// Init returns the init, if any
func (obj *execution) Init() inits.Init {
	return obj.init
}

// IsBegin returns true if there is a begin, false otherwise
func (obj *execution) IsBegin() bool {
	return obj.begin != nil
}

// Begin returns the begin, if any
func (obj *execution) Begin() begins.Begin {
	return obj.begin
}

// IsExecute returns true if there is an execute, false otherwise
func (obj *execution) IsExecute() bool {
	return obj.execute != nil
}

// Execute returns the execute, if any
func (obj *execution) Execute() executes.Execute {
	return obj.execute
}

// IsRetrieve returns true if there is a retrieve, false otherwise
func (obj *execution) IsRetrieve() bool {
	return obj.retrieve != nil
}

// Retrieve returns the retrieve, if any
func (obj *execution) Retrieve() retrieves.Retrieve {
	return obj.retrieve
}

// IsAmount returns true if there is an amount, false otherwise
func (obj *execution) IsAmount() bool {
	return obj.amount != nil
}

// Amount returns the amount, if any
func (obj *execution) Amount() amounts.Amount {
	return obj.amount
}

// IsHead returns true if there is an head, false otherwise
func (obj *execution) IsHead() bool {
	return obj.head != nil
}

// Head returns the head, if any
func (obj *execution) Head() heads.Head {
	return obj.head
}
