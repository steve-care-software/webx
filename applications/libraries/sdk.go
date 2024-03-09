package libraries

import (
	"github.com/steve-care-software/datastencil/applications/libraries/compilers"
	"github.com/steve-care-software/datastencil/applications/libraries/databases"
	"github.com/steve-care-software/datastencil/applications/libraries/executions"
)

// Application represents the library application
type Application interface {
	Compiler() compilers.Application
	Execution() executions.Application
	Database() databases.Application
}
