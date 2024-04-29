package updates

import (
	"github.com/steve-care-software/datastencil/domain/accounts"
	"github.com/steve-care-software/datastencil/domain/accounts/criterias"
	"github.com/steve-care-software/datastencil/domain/accounts/encryptors"
	"github.com/steve-care-software/datastencil/domain/accounts/signers"
	"github.com/steve-care-software/datastencil/domain/instances/links/elements/layers/instructions/accounts/updates"
	"github.com/steve-care-software/datastencil/domain/stacks"
)

// NewApplication creates a new application
func NewApplication(
	repository accounts.Repository,
	service accounts.Service,
	bitRate int,
) Application {
	accountBuilder := accounts.NewBuilder()
	criteriaBuilder := criterias.NewBuilder()
	signerFactory := signers.NewFactory()
	encryptorBuilder := encryptors.NewBuilder()
	return createApplication(
		repository,
		service,
		accountBuilder,
		criteriaBuilder,
		signerFactory,
		encryptorBuilder,
		bitRate,
	)
}

// Application represents an execution account application
type Application interface {
	Execute(frame stacks.Frame, instruction updates.Update) (*uint, error)
}
