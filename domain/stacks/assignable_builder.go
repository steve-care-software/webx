package stacks

import (
	"errors"

	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/stacks/accounts"
	"github.com/steve-care-software/datastencil/domain/stacks/databases"
	"github.com/steve-care-software/datastencil/domain/stacks/libraries"
)

type assignableBuilder struct {
	pBool *bool
	bytes []byte
	hash  hash.Hash
}

func createAssignableBuilder() AssignableBuilder {
	out := assignableBuilder{
		pBool: nil,
		bytes: nil,
		hash:  nil,
	}

	return &out
}

// Create initializes the builder
func (app *assignableBuilder) Create() AssignableBuilder {
	return createAssignableBuilder()
}

// WithBool adds a bool to the builder
func (app *assignableBuilder) WithBool(boolValue bool) AssignableBuilder {
	app.pBool = &boolValue
	return app
}

// WithBytes add bytes to the builder
func (app *assignableBuilder) WithBytes(bytes []byte) AssignableBuilder {
	app.bytes = bytes
	return app
}

// WithHash add hash to the builder
func (app *assignableBuilder) WithHash(hash hash.Hash) AssignableBuilder {
	app.hash = hash
	return app
}

// WithHashList adds an hash list to the builder
func (app *assignableBuilder) WithHashList(hashList []hash.Hash) AssignableBuilder {
	return nil
}

// WithStringList adds a string list to the builder
func (app *assignableBuilder) WithStringList(strList []string) AssignableBuilder {
	return nil
}

// WithUnsignedInt adds an uint to the builder
func (app *assignableBuilder) WithUnsignedInt(unsignedInt uint) AssignableBuilder {
	return nil
}

// WithAccount adds an account to the builder
func (app *assignableBuilder) WithAccount(account accounts.Account) AssignableBuilder {
	return nil
}

// WithLibrary adds a library to the builder
func (app *assignableBuilder) WithLibrary(lib libraries.Library) AssignableBuilder {
	return nil
}

// WithDatabase adds a database to the builder
func (app *assignableBuilder) WithDatabase(database databases.Database) AssignableBuilder {
	return nil
}

// Now builds a new Assignable instance
func (app *assignableBuilder) Now() (Assignable, error) {
	if app.pBool != nil {
		return createAssignableWithBool(app.pBool), nil
	}

	if app.bytes != nil && len(app.bytes) <= 0 {
		app.bytes = nil
	}

	if app.bytes != nil {
		return createAssignableWithBytes(app.bytes), nil
	}

	if app.hash != nil {
		return createAssignableWithHash(app.hash), nil
	}

	return nil, errors.New("the Assignable is invalid")
}
