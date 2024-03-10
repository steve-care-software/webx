package layers

import (
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/libraries/layers/instructions"
	"github.com/steve-care-software/datastencil/domain/instances/libraries/layers/outputs"
)

type layer struct {
	hash         hash.Hash
	instructions instructions.Instructions
	output       outputs.Output
	input        string
}

func createLayer(
	hash hash.Hash,
	instructions instructions.Instructions,
	output outputs.Output,
	input string,
) Layer {
	return createLayerInternally(
		hash,
		instructions,
		output,
		input,
	)
}

func createLayerInternally(
	hash hash.Hash,
	instructions instructions.Instructions,
	output outputs.Output,
	input string,
) Layer {
	out := layer{
		hash:         hash,
		instructions: instructions,
		output:       output,
		input:        input,
	}

	return &out
}

// Hash returns the hash
func (obj *layer) Hash() hash.Hash {
	return obj.hash
}

// Instructions returns the instructions
func (obj *layer) Instructions() instructions.Instructions {
	return obj.instructions
}

// Output returns the output
func (obj *layer) Output() outputs.Output {
	return obj.output
}

// Input returns the input
func (obj *layer) Input() string {
	return obj.input
}
