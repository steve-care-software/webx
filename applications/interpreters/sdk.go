package interpreters

import (
	"github.com/steve-care-software/webx/domain/cryptography/hash"
	"github.com/steve-care-software/webx/domain/programs"
)

// NewApplication creates a new interpreter application
func NewApplication() Application {
	hashAdapter := hash.NewAdapter()
	return createApplication(hashAdapter)
}

// Application represents an interpreter application
type Application interface {
	Execute(input map[string]interface{}, program programs.Program) (map[string]interface{}, error)
}
