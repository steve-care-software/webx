package references

import (
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances"
)

type reference struct {
	hash     hash.Hash
	variable string
	instance instances.Instance
}

func createReference(
	hash hash.Hash,
	variable string,
	instance instances.Instance,
) Reference {
	out := reference{
		hash:     hash,
		variable: variable,
		instance: instance,
	}

	return &out
}

// Hash returns the hash
func (obj *reference) Hash() hash.Hash {
	return obj.hash
}

// Variable returns the variable
func (obj *reference) Variable() string {
	return obj.variable
}

// Instance returns the instance
func (obj *reference) Instance() instances.Instance {
	return obj.instance
}
