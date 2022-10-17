package identities

import "github.com/steve-care-software/syntax/applications/engines/identities/authenticates"

// Application represents an identity application
type Application interface {
	List() ([]string, error)
	Select(name string) (authenticates.Application, error)
}
