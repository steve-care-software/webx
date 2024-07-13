package layers

import (
	"github.com/steve-care-software/datastencil/domain/instances/layers/instructions"
	"github.com/steve-care-software/datastencil/domain/instances/layers/outputs"
	"github.com/steve-care-software/datastencil/domain/instances/layers/references"
	"github.com/steve-care-software/historydb/domain/hash"
)

type layer struct {
	hash         hash.Hash
	instructions instructions.Instructions
	output       outputs.Output
	input        string
	references   references.References
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
		nil,
	)
}

func createLayerWithReferences(
	hash hash.Hash,
	instructions instructions.Instructions,
	output outputs.Output,
	input string,
	references references.References,
) Layer {
	return createLayerInternally(
		hash,
		instructions,
		output,
		input,
		references,
	)
}

func createLayerInternally(
	hash hash.Hash,
	instructions instructions.Instructions,
	output outputs.Output,
	input string,
	references references.References,
) Layer {
	out := layer{
		hash:         hash,
		instructions: instructions,
		output:       output,
		input:        input,
		references:   references,
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

// HasReferences returns true if there is references, false otherwise
func (obj *layer) HasReferences() bool {
	return obj.references != nil
}

// References returns the references, if any
func (obj *layer) References() references.References {
	return obj.references
}
