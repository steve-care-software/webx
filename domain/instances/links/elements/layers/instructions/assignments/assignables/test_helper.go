package assignables

import (
	"github.com/steve-care-software/datastencil/domain/instances/links/elements/layers/instructions/assignments/assignables/bytes"
	"github.com/steve-care-software/datastencil/domain/instances/links/elements/layers/instructions/assignments/assignables/compilers"
	"github.com/steve-care-software/datastencil/domain/instances/links/elements/layers/instructions/assignments/assignables/constants"
	"github.com/steve-care-software/datastencil/domain/instances/links/elements/layers/instructions/assignments/assignables/cryptography"
)

// NewAssignableWithBytesForTests creates a new assignable with bytes for tests
func NewAssignableWithBytesForTests(input bytes.Bytes) Assignable {
	ins, err := NewBuilder().Create().WithBytes(input).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewAssignableWithConstantForTests creates a new assignable with constant for tests
func NewAssignableWithConstantForTests(constant constants.Constant) Assignable {
	ins, err := NewBuilder().Create().WithConsant(constant).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewAssignableWithCryptographyForTests creates a new assignable with cryptography for tests
func NewAssignableWithCryptographyForTests(cryptography cryptography.Cryptography) Assignable {
	ins, err := NewBuilder().Create().WithCryptography(cryptography).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewAssignableWithCompilerForTests creates a new assignable with compiler for tests
func NewAssignableWithCompilerForTests(compiler compilers.Compiler) Assignable {
	ins, err := NewBuilder().Create().WithCompiler(compiler).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
