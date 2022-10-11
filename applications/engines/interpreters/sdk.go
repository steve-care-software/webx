package interpreters

import (
	"github.com/steve-care-software/syntax/domain/syntax/outputs"
	"github.com/steve-care-software/syntax/domain/syntax/programs"
)

// NewApplication creates a new interpreter application
func NewApplication() Application {
	builder := outputs.NewBuilder()
	variableBuilder := outputs.NewVariableBuilder()
	return createApplication(builder, variableBuilder)
}

// Application represents an interpreter application
type Application interface {
	Execute(input map[string]interface{}, program programs.Program) (map[string]interface{}, error)
}
