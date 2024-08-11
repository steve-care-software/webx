package headers

import "github.com/steve-care-software/webx/engine/cursors/domain/loaders/namespaces/singles/versions/singles/workspaces/singles/branches/singles/states/singles/pointers/storages"

type header struct {
	identities storages.Storage
}

func createHeaderWithIdentities(
	identities storages.Storage,
) Header {
	return createHeaderInternally(identities)
}

func createHeaderInternally(
	identities storages.Storage,
) Header {
	out := header{
		identities: identities,
	}

	return &out
}

// HasIdentities returns true if there is identities, false otherwise
func (obj *header) HasIdentities() bool {
	return obj.identities != nil
}

// Identities returns the identities, if any
func (obj *header) Identities() storages.Storage {
	return obj.identities
}
