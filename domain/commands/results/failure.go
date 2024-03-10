package results

import "github.com/steve-care-software/datastencil/domain/hash"

type failure struct {
	hash            hash.Hash
	index           uint
	code            uint
	isRaisedInLayer bool
}

func createFailure(
	hash hash.Hash,
	index uint,
	code uint,
	isRaisedInLayer bool,
) Failure {
	return createFailureInternally(hash, index, code, isRaisedInLayer)
}

func createFailureInternally(
	hash hash.Hash,
	index uint,
	code uint,
	isRaisedInLayer bool,
) Failure {
	out := failure{
		hash:            hash,
		index:           index,
		code:            code,
		isRaisedInLayer: isRaisedInLayer,
	}

	return &out
}

// Hash returns the hash
func (obj *failure) Hash() hash.Hash {
	return obj.hash
}

// Index returns the index
func (obj *failure) Index() uint {
	return obj.index
}

// Code returns the code
func (obj *failure) Code() uint {
	return obj.code
}

// IsRaisedInLayer returns true if raisedInLayer, false otherwise
func (obj *failure) IsRaisedInLayer() bool {
	return obj.isRaisedInLayer
}
