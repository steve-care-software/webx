package databases

import (
	"github.com/steve-care-software/webx/applications/databases/blockchains"
	"github.com/steve-care-software/webx/applications/databases/connections"
	"github.com/steve-care-software/webx/applications/databases/heads"
	"github.com/steve-care-software/webx/applications/databases/identities"
	"github.com/steve-care-software/webx/applications/databases/pendings"
)

// Application represents a database application
type Application interface {
	List() ([]string, error)
	Head(name string) heads.Application
	Pendings(name string) pendings.Application
	Connection(name string) connections.Application
	Blockchain(name string) blockchains.Application
	Identity(name string) identities.Application
}
