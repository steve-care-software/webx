package fetchers

import (
	"github.com/steve-care-software/webx/databases/domain/cryptography/hash"
)

type content struct {
	pRecursive *hash.Hash
	pSelector  *hash.Hash
}

func createContentWithRecursive(
	pRecursive *hash.Hash,
) Content {
	return createContentInternally(pRecursive, nil)
}

func createContentWithSelector(
	pSelector *hash.Hash,
) Content {
	return createContentInternally(nil, pSelector)
}

func createContentInternally(
	pRecursive *hash.Hash,
	pSelector *hash.Hash,
) Content {
	out := content{
		pRecursive: pRecursive,
		pSelector:  pSelector,
	}

	return &out
}

// IsRecursive returns true if recursive, false otherwise
func (obj *content) IsRecursive() bool {
	return obj.pRecursive != nil
}

// Recursive returns the recursive, if any
func (obj *content) Recursive() *hash.Hash {
	return obj.pRecursive
}

// IsSelector returns true if selector, false otherwise
func (obj *content) IsSelector() bool {
	return obj.pSelector != nil
}

// Selector returns the selector, if any
func (obj *content) Selector() *hash.Hash {
	return obj.pSelector
}
