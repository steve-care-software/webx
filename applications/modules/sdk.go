package modules

import (
	"github.com/steve-care-software/syntax/applications/modules/identity"
	"github.com/steve-care-software/syntax/applications/modules/syntax"
)

// Application represents the action application
type Application interface {
	Syntax() syntax.Application
	Identity() identity.Application
}
