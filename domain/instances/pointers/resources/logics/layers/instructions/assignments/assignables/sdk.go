package assignables

import (
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/pointers/resources/logics/layers/instructions/assignments/assignables/bytes"
	"github.com/steve-care-software/datastencil/domain/instances/pointers/resources/logics/layers/instructions/assignments/assignables/compilers"
	"github.com/steve-care-software/datastencil/domain/instances/pointers/resources/logics/layers/instructions/assignments/assignables/constants"
	"github.com/steve-care-software/datastencil/domain/instances/pointers/resources/logics/layers/instructions/assignments/assignables/cryptography"
	"github.com/steve-care-software/datastencil/domain/instances/pointers/resources/logics/layers/instructions/assignments/assignables/databases"
	"github.com/steve-care-software/datastencil/domain/instances/pointers/resources/logics/layers/instructions/assignments/assignables/lists"
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
	WithDatabase(database databases.Database) Builder
	WithList(list lists.List) Builder
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
	IsDatabase() bool
	Database() databases.Database
	IsList() bool
	List() lists.List
}
