package executions

import (
	"github.com/steve-care-software/webx/engine/states/domain/hash"
	"github.com/steve-care-software/webx/engine/stencils/domain/instances/layers/instructions/assignments/assignables/executions/executes"
	"github.com/steve-care-software/webx/engine/stencils/domain/instances/layers/instructions/assignments/assignables/executions/inits"
	"github.com/steve-care-software/webx/engine/stencils/domain/instances/layers/instructions/assignments/assignables/executions/retrieves"
)

type content struct {
	hash     hash.Hash
	isList   bool
	init     inits.Init
	begin    string
	execute  executes.Execute
	retrieve retrieves.Retrieve
	amount   string
	head     string
}

func createContentWithList(
	hash hash.Hash,
) Content {
	return createContentInternally(hash, true, nil, "", nil, nil, "", "")
}

func createContentWithInit(
	hash hash.Hash,
	init inits.Init,
) Content {
	return createContentInternally(hash, false, init, "", nil, nil, "", "")
}

func createContentWithBegin(
	hash hash.Hash,
	begin string,
) Content {
	return createContentInternally(hash, false, nil, begin, nil, nil, "", "")
}

func createContentWithExecute(
	hash hash.Hash,
	execute executes.Execute,
) Content {
	return createContentInternally(hash, false, nil, "", execute, nil, "", "")
}

func createContentWithRetrieve(
	hash hash.Hash,
	retrieve retrieves.Retrieve,
) Content {
	return createContentInternally(hash, false, nil, "", nil, retrieve, "", "")
}

func createContentWithAmount(
	hash hash.Hash,
	amount string,
) Content {
	return createContentInternally(hash, false, nil, "", nil, nil, amount, "")
}

func createContentWithHead(
	hash hash.Hash,
	head string,
) Content {
	return createContentInternally(hash, false, nil, "", nil, nil, "", head)
}

func createContentInternally(
	hash hash.Hash,
	isList bool,
	init inits.Init,
	begin string,
	execute executes.Execute,
	retrieve retrieves.Retrieve,
	amount string,
	head string,
) Content {
	out := content{
		hash:     hash,
		isList:   isList,
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
	return obj.isList
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
	return obj.begin != ""
}

// Begin returns the begin, if any
func (obj *content) Begin() string {
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
	return obj.amount != ""
}

// Amount returns the amount, if any
func (obj *content) Amount() string {
	return obj.amount
}

// IsHead returns true if there is an head, false otherwise
func (obj *content) IsHead() bool {
	return obj.head != ""
}

// Head returns the head, if any
func (obj *content) Head() string {
	return obj.head
}
