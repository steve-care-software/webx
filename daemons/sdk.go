package daemons

import "github.com/steve-care-software/syntax/domain/identity"

// Builder represents a daemon application
type Builder interface {
	Create() Builder
	WithIdentity(identity identity.Identity) Builder
	Now() (Daemon, error)
}

// Daemon represents a daemon
type Daemon interface {
	Start()
	Stop()
}
