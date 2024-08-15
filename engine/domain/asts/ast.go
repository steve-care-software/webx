package asts

import "github.com/steve-care-software/webx/engine/domain/hash"

type ast struct {
	library    NFTs
	entry      hash.Hash
	complexity map[string]uint
}

func createAST(
	library NFTs,
	entry hash.Hash,
	complexity map[string]uint,
) AST {
	out := ast{
		library:    library,
		entry:      entry,
		complexity: complexity,
	}

	return &out
}

// Library returns the library
func (obj *ast) Library() NFTs {
	return obj.library
}

// Entry returns the entry hash
func (obj *ast) Entry() hash.Hash {
	return obj.entry
}

// Complexity returns the complexity
func (obj *ast) Complexity() map[string]uint {
	return obj.complexity
}
