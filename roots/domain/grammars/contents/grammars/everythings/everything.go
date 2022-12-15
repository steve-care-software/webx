package everythings

import (
	"github.com/steve-care-software/webx/roots/domain/blockchains/cryptography/hash"
)

type everyting struct {
	hash      hash.Hash
	exception hash.Hash
	pEscape   *hash.Hash
}

func createEverything(
	hash hash.Hash,
	exception hash.Hash,
) Everything {
	return createEverythingInternally(hash, exception, nil)
}

func createEverythingWithEscape(
	hash hash.Hash,
	exception hash.Hash,
	escape *hash.Hash,
) Everything {
	return createEverythingInternally(hash, exception, escape)
}

func createEverythingInternally(
	hash hash.Hash,
	exception hash.Hash,
	escape *hash.Hash,
) Everything {
	out := everyting{
		hash:      hash,
		exception: exception,
		pEscape:   escape,
	}

	return &out
}

// Hash returns the hash
func (obj *everyting) Hash() hash.Hash {
	return obj.hash
}

// Exception returns the exception
func (obj *everyting) Exception() hash.Hash {
	return obj.exception
}

// HasEscape returns true if there is an escape, false otherwise
func (obj *everyting) HasEscape() bool {
	return obj.pEscape != nil
}

// Escape returns the escape, if any
func (obj *everyting) Escape() *hash.Hash {
	return obj.pEscape
}
