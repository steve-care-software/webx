package identity

import (
	"github.com/steve-care-software/syntax/applications/actions/identity/authenticates"
	"github.com/steve-care-software/syntax/domain/identity/identities/publics"
)

// Application represents an application
type Application interface {
	New(name string, password []byte) error
	Authenticate(name string, password []byte) (authenticates.Application, error)
	Connections() (publics.Publics, error)
	Identity() (publics.Public, error)
}
