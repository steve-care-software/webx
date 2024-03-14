package assignables

import (
	"github.com/steve-care-software/datastencil/domain/instances/links/layers/instructions/assignments/assignables/accounts"
	"github.com/steve-care-software/datastencil/domain/instances/links/layers/instructions/assignments/assignables/bytes"
	"github.com/steve-care-software/datastencil/domain/instances/links/layers/instructions/assignments/assignables/compilers"
	"github.com/steve-care-software/datastencil/domain/instances/links/layers/instructions/assignments/assignables/constants"
	"github.com/steve-care-software/datastencil/domain/instances/links/layers/instructions/assignments/assignables/cryptography"
	"github.com/steve-care-software/datastencil/domain/instances/links/layers/instructions/assignments/assignables/databases"
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

// NewAssignableWithAccountForTests creates a new assignable with account for tests
func NewAssignableWithAccountForTests(account accounts.Account) Assignable {
	ins, err := NewBuilder().Create().WithAccount(account).Now()
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

// NewAssignableWithDatabaseForTests creates a new assignable with database for tests
func NewAssignableWithDatabaseForTests(database databases.Database) Assignable {
	ins, err := NewBuilder().Create().WithDatabase(database).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewAssignableWithQueryForTests creates a new assignable with query for tests
func NewAssignableWithQueryForTests(query string) Assignable {
	ins, err := NewBuilder().Create().WithQuery(query).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
