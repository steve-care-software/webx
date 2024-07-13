package pointers

import (
	"github.com/steve-care-software/historydb/domain/hash"
)

type pointers struct {
	hash hash.Hash
	list []Pointer
}

func createPointers(
	hash hash.Hash,
	list []Pointer,
) Pointers {
	out := pointers{
		hash: hash,
		list: list,
	}

	return &out
}

// Hash returns the hash
func (obj *pointers) Hash() hash.Hash {
	return obj.hash
}

// List returns the list
func (obj *pointers) List() []Pointer {
	return obj.list
}

// First returns the first pointer
func (obj *pointers) First() Pointer {
	return obj.list[0]
}

// Match returns the matched pointers
func (obj *pointers) Match(executed [][]string) []Pointer {
	output := []Pointer{}
	for _, onePointer := range obj.list {
		// if not active, just skip it:
		if !onePointer.IsActive() {
			continue
		}

		if onePointer.HasCanceller() {
			canceller := onePointer.Canceller()
			if canceller.Match(executed) {
				continue
			}
		}

		if onePointer.HasLoader() {
			loader := onePointer.Loader()
			if loader.Match(executed) {
				output = append(output, onePointer)
				continue
			}
		}

		// no canceller, no loader:
		output = append(output, onePointer)
	}

	return output
}
