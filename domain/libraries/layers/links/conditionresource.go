package links

import "github.com/steve-care-software/identity/domain/hash"

type conditionResource struct {
	hash           hash.Hash
	code           uint
	isRaiseInLayer bool
}

func createConditionResource(
	hash hash.Hash,
	code uint,
	isRaiseInLayer bool,
) ConditionResource {
	out := conditionResource{
		hash:           hash,
		code:           code,
		isRaiseInLayer: isRaiseInLayer,
	}

	return &out
}

// Hash returns the hash
func (obj *conditionResource) Hash() hash.Hash {
	return obj.hash
}

// Code returns the code
func (obj *conditionResource) Code() uint {
	return obj.code
}

// IsRaisedInLayer returns true if raisedInLayer, false otherwise
func (obj *conditionResource) IsRaisedInLayer() bool {
	return obj.isRaiseInLayer
}
