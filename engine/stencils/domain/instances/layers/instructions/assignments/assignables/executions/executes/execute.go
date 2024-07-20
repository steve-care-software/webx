package executes

import (
	"github.com/steve-care-software/webx/engine/states/domain/hash"
	"github.com/steve-care-software/webx/engine/stencils/domain/instances/layers/instructions/assignments/assignables/executions/executes/inputs"
)

type execute struct {
	hash    hash.Hash
	context string
	input   inputs.Input
	layer   string
}

func createExecute(
	hash hash.Hash,
	context string,
	input inputs.Input,
) Execute {
	return createExecuteInternally(hash, context, input, "")
}

func createExecuteWithLayer(
	hash hash.Hash,
	context string,
	input inputs.Input,
	layer string,
) Execute {
	return createExecuteInternally(hash, context, input, layer)
}

func createExecuteInternally(
	hash hash.Hash,
	context string,
	input inputs.Input,
	layer string,
) Execute {
	out := execute{
		hash:    hash,
		context: context,
		input:   input,
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

// HasLayer returns true if there is a layer, false otherwise
func (obj *execute) HasLayer() bool {
	return obj.layer != ""
}

// Layer returns the layer, if any
func (obj *execute) Layer() string {
	return obj.layer
}
