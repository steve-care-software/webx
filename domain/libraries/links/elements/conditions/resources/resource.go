package resources

import "github.com/steve-care-software/datastencil/domain/hash"

type resource struct {
	hash           hash.Hash
	code           uint
	isRaiseInLayer bool
}

func createResource(
	hash hash.Hash,
	code uint,
	isRaiseInLayer bool,
) Resource {
	out := resource{
		hash:           hash,
		code:           code,
		isRaiseInLayer: isRaiseInLayer,
	}

	return &out
}

// Hash returns the hash
func (obj *resource) Hash() hash.Hash {
	return obj.hash
}

// Code returns the code
func (obj *resource) Code() uint {
	return obj.code
}

// IsRaisedInLayer returns true if raisedInLayer, false otherwise
func (obj *resource) IsRaisedInLayer() bool {
	return obj.isRaiseInLayer
}
