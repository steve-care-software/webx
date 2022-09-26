package modules

import (
	"github.com/steve-care-software/syntax/applications/modules/methods"
	"github.com/steve-care-software/syntax/applications/modules/objects"
	"github.com/steve-care-software/syntax/domain/programs/instructions/applications/modules"
)

// Application represents the module application
type Application interface {
	All() modules.Modules
	Methods() methods.Application
	Objects() objects.Application
}
