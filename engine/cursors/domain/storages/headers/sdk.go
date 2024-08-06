package headers

import (
	"github.com/steve-care-software/webx/engine/cursors/domain/storages/headers/blockchains"
	"github.com/steve-care-software/webx/engine/cursors/domain/storages/headers/identities"
	"github.com/steve-care-software/webx/engine/cursors/domain/storages/headers/namespaces"
)

// Header represents an header
type Header interface {
	HasIdentities() bool
	Identities() identities.Identities
	HasNamespaces() bool
	Namespaces() namespaces.Namespaces
	HasBlockchain() bool
	Blockchain() blockchains.Blockchain
}
