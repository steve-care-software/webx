package loaders

import "github.com/steve-care-software/webx/engine/cursors/domain/loaders/identities"

type loader struct {
	identity identities.Identity
}

func createLoader() Loader {
	return createLoaderInternally(nil)
}

func createLoaderWithIdentity(
	identity identities.Identity,
) Loader {
	return createLoaderInternally(identity)
}

func createLoaderInternally(
	identity identities.Identity,
) Loader {
	out := loader{
		identity: identity,
	}

	return &out
}

// NextIndex returns the next pointer index
func (obj *loader) NextIndex() (*uint64, error) {
	return nil, nil
}

// HasIdentity returns true if there is an identity, false otherwise
func (obj *loader) HasIdentity() bool {
	return obj.identity != nil
}

// Identity returns the identity, if any
func (obj *loader) Identity() identities.Identity {
	return obj.identity
}
