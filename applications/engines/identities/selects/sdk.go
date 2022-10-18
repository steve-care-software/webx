package selects

import (
	"github.com/steve-care-software/syntax/domain/syntax/identities"
	"github.com/steve-care-software/syntax/domain/syntax/identities/modifications"
)

// NewBuilder creates a new builder instance
func NewBuilder(
	repository identities.Repository,
	service identities.Service,
) Builder {
	builder := identities.NewBuilder()
	modificationsBuilder := modifications.NewBuilder()
	return createBuilder(
		builder,
		modificationsBuilder,
		repository,
		service,
	)
}

// Builder represents the application builder
type Builder interface {
	Create() Builder
	WithName(name string) Builder
	Now() (Application, error)
}

// Application represents the selected identity application
type Application interface {
	Retrieve(password string) (identities.Identity, error)
	Modify(modification modifications.Modification, currentPassword string, newPassword string) error
}
