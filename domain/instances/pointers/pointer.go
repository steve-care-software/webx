package pointers

import (
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/pointers/conditions"
)

type pointer struct {
	hash      hash.Hash
	path      []string
	isActive  bool
	loader    conditions.Condition
	canceller conditions.Condition
}

func createPointer(
	hash hash.Hash,
	path []string,
	isActive bool,
) Pointer {
	return createPointerInternally(hash, path, isActive, nil, nil)
}

func createPointerWithLoader(
	hash hash.Hash,
	path []string,
	isActive bool,
	loader conditions.Condition,
) Pointer {
	return createPointerInternally(hash, path, isActive, loader, nil)
}

func createPointerWithCanceller(
	hash hash.Hash,
	path []string,
	isActive bool,
	canceller conditions.Condition,
) Pointer {
	return createPointerInternally(hash, path, isActive, nil, canceller)
}

func createPointerWithLoaderAndCanceller(
	hash hash.Hash,
	path []string,
	isActive bool,
	loader conditions.Condition,
	canceller conditions.Condition,
) Pointer {
	return createPointerInternally(hash, path, isActive, loader, canceller)
}

func createPointerInternally(
	hash hash.Hash,
	path []string,
	isActive bool,
	loader conditions.Condition,
	canceller conditions.Condition,
) Pointer {
	out := pointer{
		hash:      hash,
		path:      path,
		isActive:  isActive,
		loader:    loader,
		canceller: canceller,
	}

	return &out
}

// Hash returns the hash
func (obj *pointer) Hash() hash.Hash {
	return obj.hash
}

// Path returns the path
func (obj *pointer) Path() []string {
	return obj.path
}

// IsActive returns true if active, false otherwise
func (obj *pointer) IsActive() bool {
	return obj.isActive
}

// HasLoader returns true if there is a loading condition, false otherwise
func (obj *pointer) HasLoader() bool {
	return obj.loader != nil
}

// Loader returns the loading condition, if any
func (obj *pointer) Loader() conditions.Condition {
	return obj.loader
}

// HasCanceller returns true if there is a cancel condition, false otherwise
func (obj *pointer) HasCanceller() bool {
	return obj.canceller != nil
}

// Canceller returns the cancel condition, if any
func (obj *pointer) Canceller() conditions.Condition {
	return obj.canceller
}
