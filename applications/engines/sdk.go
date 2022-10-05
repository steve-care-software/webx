package engines

import (
	"github.com/steve-care-software/syntax/applications/engines/criterias"
	"github.com/steve-care-software/syntax/applications/engines/grammars"
	"github.com/steve-care-software/syntax/applications/engines/interpreters"
)

// Builder represents an application builder
type Builder interface {
	Create() Builder
	WithInterpreter(interpreter interpreters.Application) Builder
	Now() (Application, error)
}

// Application represents an engine application
type Application interface {
	Grammar() grammars.Application
	Criteria() criterias.Application
	Interpreter() interpreters.Application
}
