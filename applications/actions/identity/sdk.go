package identity

import (
	"github.com/steve-care-software/syntax/applications/actions/identity/authenticates"
	"github.com/steve-care-software/syntax/applications/actions/identity/publics"
)

// Application represents an application
type Application interface {
	List() ([]string, error)
	New(name string, password []byte) error
	Public(name string) (publics.Application, error)
	Authenticate(name string, password []byte) (authenticates.Application, error)
}
