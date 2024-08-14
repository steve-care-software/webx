package identities

import (
	"github.com/steve-care-software/webx/engine/domain/identities"
	"github.com/steve-care-software/webx/engine/domain/identities/profiles"
)

// Application represents the identity application
type Application interface {
	List() ([]string, error)                                                                               // list the name of the identities
	Create(profile profiles.Profile) (identities.Identity, error)                                          // creates a new identity
	Insert(identity identities.Identity, password []byte) error                                            // inserts an identity
	Update(name string, update identities.Identity, originalPassword []byte, updatedPassword []byte) error // updates an identity
	Delete(name string, password []byte) error                                                             // deletes an identity
	Authenticate(name string, password []byte) (identities.Identity, error)                                // authenticate
}
