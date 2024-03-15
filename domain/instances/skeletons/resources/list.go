package resources

import "github.com/steve-care-software/datastencil/domain/hash"

type list struct {
	hash      hash.Hash
	value     uint8
	delimiter string
}

func createList(
	hash hash.Hash,
	value uint8,
	delimiter string,
) List {
	out := list{
		hash:      hash,
		value:     value,
		delimiter: delimiter,
	}

	return &out
}

// Hash returns the hash
func (obj *list) Hash() hash.Hash {
	return obj.hash
}

// Value returns the value
func (obj *list) Value() uint8 {
	return obj.value
}

// Delimiter returns the delimiter
func (obj *list) Delimiter() string {
	return obj.delimiter
}
