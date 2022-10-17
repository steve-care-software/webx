package grammars

import "github.com/steve-care-software/syntax/domain/syntax/databases/cryptography/hash"

type external struct {
	hash    hash.Hash
	name    string
	grammar Grammar
}

func createExternal(
	hash hash.Hash,
	name string,
	grammar Grammar,
) External {
	out := external{
		hash:    hash,
		name:    name,
		grammar: grammar,
	}

	return &out
}

// Hash returns the hash
func (obj *external) Hash() hash.Hash {
	return obj.hash
}

// Name returns the name
func (obj *external) Name() string {
	return obj.name
}

// Grammar returns the grammar
func (obj *external) Grammar() Grammar {
	return obj.grammar
}
