package modules

import "github.com/steve-care-software/syntax/domain/syntax/programs/applications/modules"

// Application represents the module application
type Application interface {
	Execute() (modules.Modules, error)
}
