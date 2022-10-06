package interpreters

import (
	"github.com/steve-care-software/syntax/applications/engines/creates"
	"github.com/steve-care-software/syntax/applications/engines/grammars"
	"github.com/steve-care-software/syntax/domain/syntax/outputs"
)

// NewApplication creates a new interpreter application
func NewApplication(
	create creates.Application,
) Application {
	grammarApp := grammars.NewApplication()
	return createApplication(nil, grammarApp, create)
}

// Application represents an interpreter application
type Application interface {
	Execute(input map[string]interface{}, script []byte) (outputs.Output, []byte, error)
}
