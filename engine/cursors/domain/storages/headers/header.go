package headers

import "github.com/steve-care-software/webx/engine/hashes/domain/pointers"

type header struct {
	identities pointers.Pointer
}

func createHeaderWithIdentities(
	identities pointers.Pointer,
) Header {
	return createHeaderInternally(identities)
}

func createHeaderInternally(
	identities pointers.Pointer,
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
func (obj *header) Identities() pointers.Pointer {
	return obj.identities
}
