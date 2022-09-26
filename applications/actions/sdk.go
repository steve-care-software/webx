package actions

import (
	"github.com/steve-care-software/syntax/applications/actions/syntax"
	"github.com/steve-care-software/syntax/applications/modules/identity"
)

// Application represents the action application
type Application interface {
	Syntax() syntax.Application
	Identity() identity.Application
}
