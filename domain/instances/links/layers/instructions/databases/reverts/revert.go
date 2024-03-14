package reverts

import "github.com/steve-care-software/datastencil/domain/hash"

type revert struct {
	hash  hash.Hash
	index string
}

func createRevert(
	hash hash.Hash,
) Revert {
	return createRevertInternally(
		hash,
		"",
	)
}

func createRevertWithIndex(
	hash hash.Hash,
	index string,
) Revert {
	return createRevertInternally(
		hash,
		index,
	)
}

func createRevertInternally(
	hash hash.Hash,
	index string,
) Revert {
	out := revert{
		hash:  hash,
		index: index,
	}

	return &out
}

// Hash returns the hash
func (obj *revert) Hash() hash.Hash {
	return obj.hash
}

// HasIndex returns true if there is an index, false otherwise
func (obj *revert) HasIndex() bool {
	return obj.index != ""
}

// Index returns the index, if any
func (obj *revert) Index() string {
	return obj.index
}
