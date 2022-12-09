package grammars

import "github.com/steve-care-software/webx/blockchains/domain/cryptography/hash"

type everything struct {
	hash      hash.Hash
	name      string
	exception Token
	escape    Token
}

func createEverything(
	hash hash.Hash,
	name string,
	exception Token,
) Everything {
	return createEverythingInternally(hash, name, exception, nil)
}

func createEverythingWithEscape(
	hash hash.Hash,
	name string,
	exception Token,
	escape Token,
) Everything {
	return createEverythingInternally(hash, name, exception, escape)
}

func createEverythingInternally(
	hash hash.Hash,
	name string,
	exception Token,
	escape Token,
) Everything {
	out := everything{
		hash:      hash,
		name:      name,
		exception: exception,
		escape:    escape,
	}

	return &out
}

// Hash returns the name
func (obj *everything) Hash() hash.Hash {
	return obj.hash
}

// Points returns the amount of points an everything contains
func (obj *everything) Points() uint {
	amount := obj.Exception().Block().Points()
	if obj.HasEscape() {
		amount += obj.Escape().Block().Points()
	}

	return amount
}

// Name returns the name
func (obj *everything) Name() string {
	return obj.name
}

// Exception returns the exception
func (obj *everything) Exception() Token {
	return obj.exception
}

// HasEscape returns true if there is an escape, false otherwise
func (obj *everything) HasEscape() bool {
	return obj.escape != nil
}

// Escape returns the escape, if any
func (obj *everything) Escape() Token {
	return obj.escape
}
