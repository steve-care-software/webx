package grammars

import "github.com/steve-care-software/webx/domain/cryptography/hash"

type suite struct {
	hash    hash.Hash
	isValid bool
	content []byte
}

func createSuiteWithValid(
	hash hash.Hash,
	valid []byte,
) Suite {
	return createSuiteInternally(hash, true, valid)
}

func createSuiteWithInvalid(
	hash hash.Hash,
	invalid []byte,
) Suite {
	return createSuiteInternally(hash, false, invalid)
}

func createSuiteInternally(
	hash hash.Hash,
	isValid bool,
	content []byte,
) Suite {
	out := suite{
		hash:    hash,
		isValid: isValid,
		content: content,
	}

	return &out
}

// Hash returns the hash
func (obj *suite) Hash() hash.Hash {
	return obj.hash
}

// IsValid returns true if valid, false otherwise
func (obj *suite) IsValid() bool {
	return obj.isValid
}

// Content returns the the content
func (obj *suite) Content() []byte {
	return obj.content
}
