package libraries

import (
	"github.com/steve-care-software/datastencil/domain/libraries/layers/instructions/assignments/assignables/libraries/compilers"
	"github.com/steve-care-software/datastencil/domain/libraries/layers/instructions/assignments/assignables/libraries/databases"
	"github.com/steve-care-software/datastencil/domain/libraries/layers/instructions/assignments/assignables/libraries/executions"
)

// Library represents a library assignable
type Library interface {
	IsCompiler() bool
	Compiler() compilers.Compiler
	IsDatabase() bool
	Database() databases.Database
	IsExecution() bool
	Execution() executions.Execution
}
