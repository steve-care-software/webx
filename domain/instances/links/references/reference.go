package references

import "github.com/steve-care-software/datastencil/domain/hash"

type reference struct {
	hash       hash.Hash
	variable   string
	identifier hash.Hash
}

func createReference(
	hash hash.Hash,
	variable string,
	identifier hash.Hash,
) Reference {
	out := reference{
		hash:       hash,
		variable:   variable,
		identifier: identifier,
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

// Identifier returns the identifier
func (obj *reference) Identifier() hash.Hash {
	return obj.identifier
}
