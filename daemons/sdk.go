package daemons

import "github.com/steve-care-software/webx/domain/identities"

// Builder represents a daemon application
type Builder interface {
	Create() Builder
	WithIdentity(identity identities.Identity) Builder
	Now() (Daemon, error)
}

// Daemon represents a daemon
type Daemon interface {
	Start()
	Stop()
}
