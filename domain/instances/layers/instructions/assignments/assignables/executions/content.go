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

type content struct {
	hash     hash.Hash
	list     string
	init     inits.Init
	begin    begins.Begin
	execute  executes.Execute
	retrieve retrieves.Retrieve
	amount   amounts.Amount
	head     heads.Head
}

func createContentWithList(
	hash hash.Hash,
	list string,
) Content {
	return createContentInternally(hash, list, nil, nil, nil, nil, nil, nil)
}

func createContentWithInit(
	hash hash.Hash,
	init inits.Init,
) Content {
	return createContentInternally(hash, "", init, nil, nil, nil, nil, nil)
}

func createContentWithBegin(
	hash hash.Hash,
	begin begins.Begin,
) Content {
	return createContentInternally(hash, "", nil, begin, nil, nil, nil, nil)
}

func createContentWithExecute(
	hash hash.Hash,
	execute executes.Execute,
) Content {
	return createContentInternally(hash, "", nil, nil, execute, nil, nil, nil)
}

func createContentWithRetrieve(
	hash hash.Hash,
	retrieve retrieves.Retrieve,
) Content {
	return createContentInternally(hash, "", nil, nil, nil, retrieve, nil, nil)
}

func createContentWithAmount(
	hash hash.Hash,
	amount amounts.Amount,
) Content {
	return createContentInternally(hash, "", nil, nil, nil, nil, amount, nil)
}

func createContentWithHead(
	hash hash.Hash,
	head heads.Head,
) Content {
	return createContentInternally(hash, "", nil, nil, nil, nil, nil, head)
}

func createContentInternally(
	hash hash.Hash,
	list string,
	init inits.Init,
	begin begins.Begin,
	execute executes.Execute,
	retrieve retrieves.Retrieve,
	amount amounts.Amount,
	head heads.Head,
) Content {
	out := content{
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
func (obj *content) Hash() hash.Hash {
	return obj.hash
}

// IsList returns true if there is a list, false otherwise
func (obj *content) IsList() bool {
	return obj.list != ""
}

// List returns the list, if any
func (obj *content) List() string {
	return obj.list
}

// IsInit returns true if there is an init, false otherwise
func (obj *content) IsInit() bool {
	return obj.init != nil
}

// Init returns the init, if any
func (obj *content) Init() inits.Init {
	return obj.init
}

// IsBegin returns true if there is a begin, false otherwise
func (obj *content) IsBegin() bool {
	return obj.begin != nil
}

// Begin returns the begin, if any
func (obj *content) Begin() begins.Begin {
	return obj.begin
}

// IsExecute returns true if there is an execute, false otherwise
func (obj *content) IsExecute() bool {
	return obj.execute != nil
}

// Execute returns the execute, if any
func (obj *content) Execute() executes.Execute {
	return obj.execute
}

// IsRetrieve returns true if there is a retrieve, false otherwise
func (obj *content) IsRetrieve() bool {
	return obj.retrieve != nil
}

// Retrieve returns the retrieve, if any
func (obj *content) Retrieve() retrieves.Retrieve {
	return obj.retrieve
}

// IsAmount returns true if there is an amount, false otherwise
func (obj *content) IsAmount() bool {
	return obj.amount != nil
}

// Amount returns the amount, if any
func (obj *content) Amount() amounts.Amount {
	return obj.amount
}

// IsHead returns true if there is an head, false otherwise
func (obj *content) IsHead() bool {
	return obj.head != nil
}

// Head returns the head, if any
func (obj *content) Head() heads.Head {
	return obj.head
}
