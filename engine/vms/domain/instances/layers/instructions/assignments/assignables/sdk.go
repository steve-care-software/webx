package assignables

import (
	"github.com/steve-care-software/webx/engine/hashes/domain/hash"
	"github.com/steve-care-software/webx/engine/vms/domain/instances/layers/instructions/assignments/assignables/bytes"
	"github.com/steve-care-software/webx/engine/vms/domain/instances/layers/instructions/assignments/assignables/compilers"
	"github.com/steve-care-software/webx/engine/vms/domain/instances/layers/instructions/assignments/assignables/constants"
	"github.com/steve-care-software/webx/engine/vms/domain/instances/layers/instructions/assignments/assignables/cryptography"
	"github.com/steve-care-software/webx/engine/vms/domain/instances/layers/instructions/assignments/assignables/executables"
	"github.com/steve-care-software/webx/engine/vms/domain/instances/layers/instructions/assignments/assignables/executions"
	"github.com/steve-care-software/webx/engine/vms/domain/instances/layers/instructions/assignments/assignables/lists"
)

// NewBuilder creates a new assignable builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// Adapter represents the assignable adapter
type Adapter interface {
	ToBytes(ins Assignable) ([]byte, error)
	ToInstance(bytes []byte) (Assignable, error)
}

// Builder represents an assignable builder
type Builder interface {
	Create() Builder
	WithBytes(bytes bytes.Bytes) Builder
	WithConsant(constant constants.Constant) Builder
	WithCryptography(cryptography cryptography.Cryptography) Builder
	WithCompiler(compiler compilers.Compiler) Builder
	WithExecution(execution executions.Execution) Builder
	WithList(list lists.List) Builder
	WithExecutable(executable executables.Executable) Builder
	WithVariable(variable string) Builder
	Now() (Assignable, error)
}

// Assignable represents an assignable
type Assignable interface {
	Hash() hash.Hash
	IsBytes() bool
	Bytes() bytes.Bytes
	IsConstant() bool
	Constant() constants.Constant
	IsCryptography() bool
	Cryptography() cryptography.Cryptography
	IsCompiler() bool
	Compiler() compilers.Compiler
	IsExecution() bool
	Execution() executions.Execution
	IsList() bool
	List() lists.List
	IsExecutable() bool
	Executable() executables.Executable
	IsVariable() bool
	Variable() string
}
