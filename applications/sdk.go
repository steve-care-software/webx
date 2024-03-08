package applications

import (
	"github.com/steve-care-software/datastencil/applications/accounts"
	"github.com/steve-care-software/datastencil/applications/compilers"
	"github.com/steve-care-software/datastencil/applications/cryptography"
	"github.com/steve-care-software/datastencil/applications/orms"
	"github.com/steve-care-software/datastencil/domain/libraries/layers/instructions/assignments/assignables/executions"
)

// Application represents the application
type Application interface {
	Account() accounts.Application
	Compile() compilers.Application
	Cryptography() cryptography.Application
	Execution() executions.Execution
	ORM() orms.Application
}
