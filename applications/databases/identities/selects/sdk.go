package selects

import (
	"github.com/steve-care-software/webx/domain/identities"
	"github.com/steve-care-software/webx/domain/identities/modifications"
)

// Builder represents the application builder
type Builder interface {
	Create() Builder
	WithName(name string) Builder
	Now() (Application, error)
}

// Application represents the selected identity application
type Application interface {
	Retrieve(password []byte) (identities.Identity, error)
	Modify(modification modifications.Modification, currentPassword []byte, newPassword []byte) error
}
