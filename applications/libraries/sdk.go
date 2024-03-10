package libraries

import (
	"github.com/steve-care-software/datastencil/applications/libraries/executions"
)

// Application represents the library application
type Application interface {
	Execution() executions.Application
}
