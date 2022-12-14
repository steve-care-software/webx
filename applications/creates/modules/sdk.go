package modules

import "github.com/steve-care-software/webx/domain/programs/modules"

// Application represents the module application
type Application interface {
	Execute() (modules.Modules, error)
}
