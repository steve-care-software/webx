package layers

import (
	"github.com/steve-care-software/webx/engine/states/domain/hash"
	"github.com/steve-care-software/webx/engine/vms/domain/instances/layers/instructions"
	"github.com/steve-care-software/webx/engine/vms/domain/instances/layers/outputs"
	"github.com/steve-care-software/webx/engine/vms/domain/instances/layers/references"
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
) Layer {
	return createLayerInternally(
		hash,
		instructions,
		output,
		"",
		nil,
	)
}

func createLayerWithReferences(
	hash hash.Hash,
	instructions instructions.Instructions,
	output outputs.Output,
	references references.References,
) Layer {
	return createLayerInternally(
		hash,
		instructions,
		output,
		"",
		references,
	)
}

func createLayerWithInput(
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

func createLayerWithReferencesAndInput(
	hash hash.Hash,
	instructions instructions.Instructions,
	output outputs.Output,
	references references.References,
	input string,
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

// HasInput returns true if there is an input, false otherwise
func (obj *layer) HasInput() bool {
	return obj.input != ""
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
