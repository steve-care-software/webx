package interpreters

import (
	"github.com/steve-care-software/syntax/domain/syntax/programs"
)

// NewApplication creates a new interpreter application
func NewApplication() Application {
	return createApplication()
}

// Application represents an interpreter application
type Application interface {
	Execute(input map[string]interface{}, program programs.Program) (map[string]interface{}, error)
}
