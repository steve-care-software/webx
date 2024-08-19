package asts

import "github.com/steve-care-software/webx/engine/domain/hash"

type ast struct {
	library NFTs
	entry   hash.Hash
}

func createAST(
	library NFTs,
	entry hash.Hash,
) AST {
	out := ast{
		library: library,
		entry:   entry,
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
