package databases

import (
	"errors"

	"github.com/steve-care-software/syntax/domain/syntax/databases/blockchains"
	"github.com/steve-care-software/syntax/domain/syntax/databases/cryptography/encryptions/keys"
	"github.com/steve-care-software/syntax/domain/syntax/databases/cryptography/hash"
	"github.com/steve-care-software/syntax/domain/syntax/databases/schemas"
)

type contentBuilder struct {
	hashAdapter      hash.Adapter
	publicKeyAdapter keys.PublicKeyAdapter
	name             string
	schema           schemas.Schema
	reference        blockchains.Blockchain
	encryptTo        keys.PublicKey
	migration        Migration
}

func createContentBuilder(
	hashAdapter hash.Adapter,
	publicKeyAdapter keys.PublicKeyAdapter,
) ContentBuilder {
	out := contentBuilder{
		hashAdapter:      hashAdapter,
		publicKeyAdapter: publicKeyAdapter,
		name:             "",
		schema:           nil,
		reference:        nil,
		encryptTo:        nil,
		migration:        nil,
	}

	return &out
}

// Create initializes the builder
func (app *contentBuilder) Create() ContentBuilder {
	return createContentBuilder(
		app.hashAdapter,
		app.publicKeyAdapter,
	)
}

// WithName adds a name to the builder
func (app *contentBuilder) WithName(name string) ContentBuilder {
	app.name = name
	return app
}

// WithSchema adds a schema to the builder
func (app *contentBuilder) WithSchema(schema schemas.Schema) ContentBuilder {
	app.schema = schema
	return app
}

// WithReference adds a Reference to the builder
func (app *contentBuilder) WithReference(reference blockchains.Blockchain) ContentBuilder {
	app.reference = reference
	return app
}

// WithMigration adds a migration to the builder
func (app *contentBuilder) WithMigration(migration Migration) ContentBuilder {
	app.migration = migration
	return app
}

// EncryptTo adds an encryption's public key to the builder
func (app *contentBuilder) EncryptTo(encryptTo keys.PublicKey) ContentBuilder {
	app.encryptTo = encryptTo
	return app
}

// Now builds a new Content instance
func (app *contentBuilder) Now() (Content, error) {
	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build a Content instance")
	}

	if app.schema == nil {
		return nil, errors.New("the schema is mandatory in order to build a Content instance")
	}

	if app.reference == nil {
		return nil, errors.New("the reference blockchain is mandatory in order to build a Content instance")
	}

	if app.encryptTo == nil {
		return nil, errors.New("the encryption's public key is mandatory in order to build a Content instance")
	}

	data := [][]byte{
		[]byte(app.name),
		app.schema.Hash().Bytes(),
		app.reference.Reference().Bytes(),
		app.publicKeyAdapter.ToBytes(app.encryptTo),
	}

	if app.migration != nil {
		data = append(data, app.migration.Hash().Bytes())
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.migration != nil {
		return createContentWithMigration(*pHash, app.name, app.schema, app.reference, app.encryptTo, app.migration), nil
	}

	return createContent(*pHash, app.name, app.schema, app.reference, app.encryptTo), nil
}
