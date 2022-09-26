package applications

import (
	"github.com/steve-care-software/syntax/applications/actions"
	"github.com/steve-care-software/syntax/applications/languages"
	"github.com/steve-care-software/syntax/applications/modules"
)

// Builder represents an application builder
type Builder interface {
	Create() Builder
	WithAction(action actions.Application) Builder
	WithLanguage(language languages.Application) Builder
	WithModule(module modules.Application) Builder
	Now() (Application, error)
}

// Application represents the syntax application
type Application interface {
	Action() actions.Application
	Language() languages.Application
	Module() modules.Application
}
