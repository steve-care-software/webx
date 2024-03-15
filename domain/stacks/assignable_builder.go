package stacks

import (
	"errors"

	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances"
	"github.com/steve-care-software/datastencil/domain/instances/queries"
	"github.com/steve-care-software/datastencil/domain/stacks/accounts"
	stack_accounts "github.com/steve-care-software/datastencil/domain/stacks/accounts"
)

type assignableBuilder struct {
	pBool        *bool
	bytes        []byte
	hash         hash.Hash
	hashList     []hash.Hash
	stringList   []string
	pUnsignedInt *uint
	account      stack_accounts.Account
	instance     instances.Instance
	query        queries.Query
}

func createAssignableBuilder() AssignableBuilder {
	out := assignableBuilder{
		pBool:        nil,
		bytes:        nil,
		hash:         nil,
		hashList:     nil,
		stringList:   nil,
		pUnsignedInt: nil,
		account:      nil,
		instance:     nil,
		query:        nil,
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
	app.hashList = hashList
	return app
}

// WithStringList adds a string list to the builder
func (app *assignableBuilder) WithStringList(strList []string) AssignableBuilder {
	app.stringList = strList
	return app
}

// WithUnsignedInt adds an uint to the builder
func (app *assignableBuilder) WithUnsignedInt(unsignedInt uint) AssignableBuilder {
	app.pUnsignedInt = &unsignedInt
	return app
}

// WithAccount adds an account to the builder
func (app *assignableBuilder) WithAccount(account accounts.Account) AssignableBuilder {
	app.account = account
	return app
}

// WithInstance adds an instance to the builder
func (app *assignableBuilder) WithInstance(instance instances.Instance) AssignableBuilder {
	app.instance = instance
	return app
}

// WithQuery adds a query to the builder
func (app *assignableBuilder) WithQuery(query queries.Query) AssignableBuilder {
	app.query = query
	return app
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

	if app.hashList != nil && len(app.hashList) <= 0 {
		app.hashList = nil
	}

	if app.hashList != nil {
		return createAssignableWithHashList(app.hashList), nil
	}

	if app.stringList != nil {
		return createAssignableWithStringList(app.stringList), nil
	}

	if app.pUnsignedInt != nil {
		return createAssignableWithUnsignedInt(app.pUnsignedInt), nil
	}

	if app.account != nil {
		return createAssignableWithAccount(app.account), nil
	}

	if app.instance != nil {
		return createAssignableWithInstance(app.instance), nil
	}

	if app.query != nil {
		return createAssignableWithQuery(app.query), nil
	}

	return nil, errors.New("the Assignable is invalid")
}
