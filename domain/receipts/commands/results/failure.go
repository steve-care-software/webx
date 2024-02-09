package results

import "github.com/steve-care-software/identity/domain/hash"

type failure struct {
	hash            hash.Hash
	code            uint
	isRaisedInLayer bool
}

func createFailure(
	hash hash.Hash,
	code uint,
	isRaisedInLayer bool,
) Failure {
	return createFailureInternally(hash, code, isRaisedInLayer)
}

func createFailureInternally(
	hash hash.Hash,
	code uint,
	isRaisedInLayer bool,
) Failure {
	out := failure{
		hash:            hash,
		code:            code,
		isRaisedInLayer: isRaisedInLayer,
	}

	return &out
}

// Hash returns the hash
func (obj *failure) Hash() hash.Hash {
	return obj.hash
}

// Code returns the code
func (obj *failure) Code() uint {
	return obj.code
}

// IsRaisedInLayer returns true if raisedInLayer, false otherwise
func (obj *failure) IsRaisedInLayer() bool {
	return obj.isRaisedInLayer
}
