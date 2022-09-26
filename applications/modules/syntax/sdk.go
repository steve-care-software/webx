package syntax

import (
	"github.com/steve-care-software/syntax/applications/modules/syntax/methods"
	"github.com/steve-care-software/syntax/applications/modules/syntax/objects"
	"github.com/steve-care-software/syntax/domain/syntax/programs/instructions/applications/modules"
)

// Application represents the module application
type Application interface {
	All() modules.Modules
	Methods() methods.Application
	Objects() objects.Application
}
