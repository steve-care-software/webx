package references

import "github.com/steve-care-software/historydb/domain/hash"

type reference struct {
	hash     hash.Hash
	variable string
	path     []string
}

func createReference(
	hash hash.Hash,
	variable string,
	path []string,
) Reference {
	out := reference{
		hash:     hash,
		variable: variable,
		path:     path,
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

// Path returns the path
func (obj *reference) Path() []string {
	return obj.path
}
