package services

import "github.com/steve-care-software/datastencil/domain/hash"

type service struct {
	hash    hash.Hash
	isBegin bool
}

func createServiceWithBegin(
	hash hash.Hash,
) Service {
	return createServiceInternally(hash, true)
}

func createServiceInternally(
	hash hash.Hash,
	isBegin bool,
) Service {
	out := service{
		hash:    hash,
		isBegin: isBegin,
	}

	return &out
}

// Hash returns the hash
func (obj *service) Hash() hash.Hash {
	return obj.hash
}

// IsBegin returns true if begin, false otherwise
func (obj *service) IsBegin() bool {
	return obj.isBegin
}
