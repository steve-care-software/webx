package omissions

import (
	"github.com/steve-care-software/webx/engine/states/domain/hash"
	"github.com/steve-care-software/webx/engine/vms/domain/instances/routes/elements"
)

type omission struct {
	hash   hash.Hash
	prefix elements.Element
	suffix elements.Element
}

func createOmissionWithPrefix(
	hash hash.Hash,
	prefix elements.Element,
) Omission {
	return createOmissionInternally(hash, prefix, nil)
}

func createOmissionWithSuffix(
	hash hash.Hash,
	suffix elements.Element,
) Omission {
	return createOmissionInternally(hash, nil, suffix)
}

func createOmissionWithPrefixAndSuffix(
	hash hash.Hash,
	prefix elements.Element,
	suffix elements.Element,
) Omission {
	return createOmissionInternally(hash, prefix, suffix)
}

func createOmissionInternally(
	hash hash.Hash,
	prefix elements.Element,
	suffix elements.Element,
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

// HasPrefix returns true if there is a prefix, false otherwise
func (obj *omission) HasPrefix() bool {
	return obj.prefix != nil
}

// Prefix returns the prefix, if any
func (obj *omission) Prefix() elements.Element {
	return obj.prefix
}

// HasSuffix returns true if there is a suffix, false otherwise
func (obj *omission) HasSuffix() bool {
	return obj.suffix != nil
}

// Suffix returns the suffix, if any
func (obj *omission) Suffix() elements.Element {
	return obj.suffix
}
