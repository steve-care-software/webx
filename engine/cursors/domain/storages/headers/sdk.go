package headers

import (
	"github.com/steve-care-software/webx/engine/cursors/domain/storages/headers/identities"
)

// Header represents an header
type Header interface {
	HasIdentities() bool
	Identities() identities.Identities
}
