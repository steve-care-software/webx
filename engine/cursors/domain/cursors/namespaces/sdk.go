package namespaces

import (
	"github.com/steve-care-software/webx/engine/cursors/domain/cursors/blockchains"
	"github.com/steve-care-software/webx/engine/cursors/domain/cursors/namespaces/versions"
)

// Namespace represents a namespace
type Namespace interface {
	Name() string
	HasVersion() bool
	Version() versions.Version
	HasBlockchain() bool
	Blockchain() blockchains.Blockchain
}
