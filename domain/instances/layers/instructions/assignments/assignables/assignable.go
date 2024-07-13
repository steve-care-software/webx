package assignables

import (
	"github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables/bytes"
	"github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables/compilers"
	"github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables/constants"
	"github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables/cryptography"
	"github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables/executions"
	"github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables/lists"
	"github.com/steve-care-software/historydb/domain/hash"
)

type assignable struct {
	hash      hash.Hash
	bytes     bytes.Bytes
	constant  constants.Constant
	crypto    cryptography.Cryptography
	compiler  compilers.Compiler
	execution executions.Execution
	list      lists.List
}

func createAssignableWithBytes(
	hash hash.Hash,
	bytes bytes.Bytes,
) Assignable {
	return createAssignableInternally(hash, bytes, nil, nil, nil, nil, nil)
}

func createAssignableWithConstant(
	hash hash.Hash,
	constant constants.Constant,
) Assignable {
	return createAssignableInternally(hash, nil, constant, nil, nil, nil, nil)
}

func createAssignableWithCryptography(
	hash hash.Hash,
	crypto cryptography.Cryptography,
) Assignable {
	return createAssignableInternally(hash, nil, nil, crypto, nil, nil, nil)
}

func createAssignableWithCompiler(
	hash hash.Hash,
	compiler compilers.Compiler,
) Assignable {
	return createAssignableInternally(hash, nil, nil, nil, compiler, nil, nil)
}

func createAssignableWithExecution(
	hash hash.Hash,
	eecution executions.Execution,
) Assignable {
	return createAssignableInternally(hash, nil, nil, nil, nil, eecution, nil)
}

func createAssignableWithList(
	hash hash.Hash,
	list lists.List,
) Assignable {
	return createAssignableInternally(hash, nil, nil, nil, nil, nil, list)
}

func createAssignableInternally(
	hash hash.Hash,
	bytes bytes.Bytes,
	constant constants.Constant,
	crypto cryptography.Cryptography,
	compiler compilers.Compiler,
	execution executions.Execution,
	list lists.List,
) Assignable {
	out := assignable{
		hash:      hash,
		bytes:     bytes,
		constant:  constant,
		crypto:    crypto,
		compiler:  compiler,
		execution: execution,
		list:      list,
	}

	return &out
}

// Hash returns the hash
func (obj *assignable) Hash() hash.Hash {
	return obj.hash
}

// IsBytes returns true if there is bytes, false otherwise
func (obj *assignable) IsBytes() bool {
	return obj.bytes != nil
}

// Bytes returns the bytes, if any
func (obj *assignable) Bytes() bytes.Bytes {
	return obj.bytes
}

// IsConstant returns true if there is constant, false otherwise
func (obj *assignable) IsConstant() bool {
	return obj.constant != nil
}

// Constant returns the constant, if any
func (obj *assignable) Constant() constants.Constant {
	return obj.constant
}

// IsCryptography returns true if there is a cryptography, false otherwise
func (obj *assignable) IsCryptography() bool {
	return obj.crypto != nil
}

// Cryptography returns the cryptography, if any
func (obj *assignable) Cryptography() cryptography.Cryptography {
	return obj.crypto
}

// IsCompiler returns true if there is a compiler, false otherwise
func (obj *assignable) IsCompiler() bool {
	return obj.compiler != nil
}

// Compiler returns the compiler, if any
func (obj *assignable) Compiler() compilers.Compiler {
	return obj.compiler
}

// IsExecution returns true if there is an execution, false otherwise
func (obj *assignable) IsExecution() bool {
	return obj.execution != nil
}

// Execution returns the execution, if any
func (obj *assignable) Execution() executions.Execution {
	return obj.execution
}

// IsList returns true if there is a list, false otherwise
func (obj *assignable) IsList() bool {
	return obj.list != nil
}

// List returns the list, if any
func (obj *assignable) List() lists.List {
	return obj.list
}
