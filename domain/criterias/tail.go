package criterias

import "github.com/steve-care-software/webx/domain/cryptography/hash"

type tail struct {
	hash      hash.Hash
	name      string
	delimiter Delimiter
}

func createTail(
	hash hash.Hash,
	name string,
) Tail {
	return createTailInternally(hash, name, nil)
}

func createTailWithDelimiter(
	hash hash.Hash,
	name string,
	delimiter Delimiter,
) Tail {
	return createTailInternally(hash, name, delimiter)
}

func createTailInternally(
	hash hash.Hash,
	name string,
	delimiter Delimiter,
) Tail {
	out := tail{
		hash:      hash,
		name:      name,
		delimiter: delimiter,
	}

	return &out
}

// Hash returns the hash
func (obj *tail) Hash() hash.Hash {
	return obj.hash
}

// Name returns the name
func (obj *tail) Name() string {
	return obj.name
}

// HasDelimiter returns true if there is a delimiter, false otherwise
func (obj *tail) HasDelimiter() bool {
	return obj.delimiter != nil
}

// Delimiter returns the delimiter, if any
func (obj *tail) Delimiter() Delimiter {
	return obj.delimiter
}
