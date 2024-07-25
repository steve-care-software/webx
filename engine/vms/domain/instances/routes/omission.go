package routes

import "github.com/steve-care-software/webx/engine/states/domain/hash"

type omission struct {
	hash   hash.Hash
	prefix Element
	suffix Element
}

func createOmissionWithPrefix(
	hash hash.Hash,
	prefix Element,
) Omission {
	return createOmissionInternally(hash, prefix, nil)
}

func createOmissionWithSuffix(
	hash hash.Hash,
	suffix Element,
) Omission {
	return createOmissionInternally(hash, nil, suffix)
}

func createOmissionWithPrefixAndSuffix(
	hash hash.Hash,
	prefix Element,
	suffix Element,
) Omission {
	return createOmissionInternally(hash, prefix, suffix)
}

func createOmissionInternally(
	hash hash.Hash,
	prefix Element,
	suffix Element,
) Omission {
	out := omission{
		hash:   hash,
		prefix: prefix,
		suffix: suffix,
	}

	return &out
}

// Hash returns the hash
func (obj *omission) Hash() hash.Hash {
	return obj.hash
}

// Remaining returns the remaining bytes
func (obj *omission) Remaining(input []byte) []byte {
	return nil
}

// HasPrefix returns true if there is a prefix, false otherwise
func (obj *omission) HasPrefix() bool {
	return obj.prefix != nil
}

// Prefix returns the prefix, if any
func (obj *omission) Prefix() Element {
	return obj.prefix
}

// HasSuffix returns true if there is a suffix, false otherwise
func (obj *omission) HasSuffix() bool {
	return obj.suffix != nil
}

// Suffix returns the suffix, if any
func (obj *omission) Suffix() Element {
	return obj.suffix
}
