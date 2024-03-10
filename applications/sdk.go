package applications

import (
	"github.com/steve-care-software/datastencil/applications/accounts"
	"github.com/steve-care-software/datastencil/applications/libraries"
)

// Application represents the application
type Application interface {
	Account() accounts.Application
	Library() libraries.Application
}
