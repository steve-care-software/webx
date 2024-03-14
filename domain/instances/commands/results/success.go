package results

import (
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/libraries/layers/outputs/kinds"
)

type success struct {
	hash   hash.Hash
	output Output
	kind   kinds.Kind
}

func createSuccess(
	hash hash.Hash,
	output Output,
	kind kinds.Kind,
) Success {
	out := success{
		hash:   hash,
		output: output,
		kind:   kind,
	}

	return &out
}

// Hash returns the hash
func (obj *success) Hash() hash.Hash {
	return obj.hash
}

// Output returns the output
func (obj *success) Output() Output {
	return obj.output
}

// Kind returns the kind
func (obj *success) Kind() kinds.Kind {
	return obj.kind
}
