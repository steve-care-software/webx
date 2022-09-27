package blockchains

import (
	"github.com/steve-care-software/syntax/applications/actions/blockchains/selects"
	"github.com/steve-care-software/syntax/domain/blockchains"
	"github.com/steve-care-software/syntax/domain/identity/cryptography/hash"
)

// NewApplication creates a new blockchain application instance
func NewApplication(
	builder selects.Builder,
	repository blockchains.Repository,
) Application {
	return createApplication(builder, repository)
}

// Application represents a blockchain application
type Application interface {
	List() ([]hash.Hash, error)
	Select(ref hash.Hash) (selects.Application, error)
}
