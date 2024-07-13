package executions

import "github.com/steve-care-software/historydb/domain/hash"

type execution struct {
	hash     hash.Hash
	isLength bool
	pFetch   *uint
}

func createExecutionWithLength(hash hash.Hash) Execution {
	return createExecutionInternally(hash, true, nil)
}

func createExecutionWithFetch(hash hash.Hash, pFetch *uint) Execution {
	return createExecutionInternally(hash, false, pFetch)
}

func createExecutionInternally(
	hash hash.Hash,
	isLength bool,
	pFetch *uint,
) Execution {
	out := execution{
		hash:     hash,
		isLength: isLength,
		pFetch:   pFetch,
	}

	return &out
}

// Hash returns the hash
func (obj *execution) Hash() hash.Hash {
	return obj.hash
}

// IsLength returns true if length, false otheriwse
func (obj *execution) IsLength() bool {
	return obj.isLength
}

// IsFetch returns true if fetch, false otherwise
func (obj *execution) IsFetch() bool {
	return obj.pFetch != nil
}

// Fetch returns the fetch if any
func (obj *execution) Fetch() *uint {
	return obj.pFetch
}
