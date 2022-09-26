package languages

import (
	"github.com/steve-care-software/syntax/applications/languages/identity"
	"github.com/steve-care-software/syntax/applications/languages/syntax"
)

// Application represents the language application
type Application interface {
	Syntax() syntax.Application
	Identity() identity.Application
}
