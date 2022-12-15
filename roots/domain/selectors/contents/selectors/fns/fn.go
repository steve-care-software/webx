package fns

import (
	"github.com/steve-care-software/webx/blockchains/domain/cryptography/hash"
)

type fn struct {
	hash      hash.Hash
	isSingle  bool
	isContent bool
	program   hash.Hash
	param     uint
}

func createFn(
	hash hash.Hash,
	isSingle bool,
	isContent bool,
	program hash.Hash,
	param uint,
) Fn {
	out := fn{
		hash:      hash,
		isSingle:  isSingle,
		isContent: isContent,
		program:   program,
		param:     param,
	}

	return &out
}

// Hash returns the hash
func (obj *fn) Hash() hash.Hash {
	return obj.hash
}

// IsSingle returns true if single, false otherwise
func (obj *fn) IsSingle() bool {
	return obj.isSingle
}

// IsContent returns true if content, false otherwise
func (obj *fn) IsContent() bool {
	return obj.isContent
}

// Program returns the program
func (obj *fn) Program() hash.Hash {
	return obj.program
}

// Param returns the param
func (obj *fn) Param() uint {
	return obj.param
}
