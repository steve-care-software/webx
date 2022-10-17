package values

import "github.com/steve-care-software/syntax/domain/syntax/databases/cryptography/hash"

type value struct {
	hash   hash.Hash
	name   string
	number byte
}

func createValue(
	hash hash.Hash,
	name string,
	number byte,
) Value {
	out := value{
		hash:   hash,
		name:   name,
		number: number,
	}

	return &out
}

// Hash returns the hash
func (obj *value) Hash() hash.Hash {
	return obj.hash
}

// Name returns the name
func (obj *value) Name() string {
	return obj.name
}

// Number returns the number
func (obj *value) Number() byte {
	return obj.number
}
