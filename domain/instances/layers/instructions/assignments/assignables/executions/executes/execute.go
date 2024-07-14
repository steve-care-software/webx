package executes

import (
	"github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables/executions/executes/inputs"
	"github.com/steve-care-software/historydb/domain/hash"
)

type execute struct {
	hash    hash.Hash
	context string
	input   inputs.Input
	ret     string
	layer   inputs.Input
}

func createExecute(
	hash hash.Hash,
	context string,
	input inputs.Input,
	ret string,
) Execute {
	return createExecuteInternally(hash, context, input, ret, nil)
}

func createExecuteWithLayer(
	hash hash.Hash,
	context string,
	input inputs.Input,
	ret string,
	layer inputs.Input,
) Execute {
	return createExecuteInternally(hash, context, input, ret, layer)
}

func createExecuteInternally(
	hash hash.Hash,
	context string,
	input inputs.Input,
	ret string,
	layer inputs.Input,
) Execute {
	out := execute{
		hash:    hash,
		context: context,
		input:   input,
		ret:     ret,
		layer:   layer,
	}

	return &out
}

// Hash returns the hash
func (obj *execute) Hash() hash.Hash {
	return obj.hash
}

// Context returns the context
func (obj *execute) Context() string {
	return obj.context
}

// Input returns the input
func (obj *execute) Input() inputs.Input {
	return obj.input
}

// Return returns the return
func (obj *execute) Return() string {
	return obj.ret
}

// HasLayer returns true if there is a layer, false otherwise
func (obj *execute) HasLayer() bool {
	return obj.layer != nil
}

// Layer returns the layer, if any
func (obj *execute) Layer() inputs.Input {
	return obj.layer
}
