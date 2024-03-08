package kinds

import "github.com/steve-care-software/datastencil/domain/hash"

type kind struct {
	hash       hash.Hash
	isPrompt   bool
	isContinue bool
}

func createKindWithPrompt(
	hash hash.Hash,
) Kind {
	return createKindInternally(hash, true, false)
}

func createKindWithContinue(
	hash hash.Hash,
) Kind {
	return createKindInternally(hash, false, true)
}

func createKindInternally(
	hash hash.Hash,
	isPrompt bool,
	isContinue bool,
) Kind {
	out := kind{
		hash:       hash,
		isPrompt:   isPrompt,
		isContinue: isContinue,
	}

	return &out
}

// Hash returns the hash
func (obj *kind) Hash() hash.Hash {
	return obj.hash
}

// IsPrompt returns true if prompt, false otherwise
func (obj *kind) IsPrompt() bool {
	return obj.isPrompt
}

// IsContinue returns true if continue, false otherwise
func (obj *kind) IsContinue() bool {
	return obj.isContinue
}
