package modifications

import (
	"errors"
	"time"

	"github.com/steve-care-software/webx/identities/domain/cryptography/encryptions/keys"
	"github.com/steve-care-software/webx/identities/domain/cryptography/signatures"
)

type modificationBuilder struct {
	name       string
	sig        signatures.PrivateKey
	enc        keys.PrivateKey
	pCreatedOn *time.Time
}

func createModificationBuilder() ModificationBuilder {
	out := modificationBuilder{
		name:       "",
		sig:        nil,
		enc:        nil,
		pCreatedOn: nil,
	}

	return &out
}

// Create initializes the builder
func (app *modificationBuilder) Create() ModificationBuilder {
	return createModificationBuilder()
}

// WithName adds a name to the builder
func (app *modificationBuilder) WithName(name string) ModificationBuilder {
	app.name = name
	return app
}

// WithSignature adds a signature pk to the builder
func (app *modificationBuilder) WithSignature(sig signatures.PrivateKey) ModificationBuilder {
	app.sig = sig
	return app
}

// WithEncryption adds an encryption pk to the builder
func (app *modificationBuilder) WithEncryption(enc keys.PrivateKey) ModificationBuilder {
	app.enc = enc
	return app
}

// CreatedOn adds a creation time to the builder
func (app *modificationBuilder) CreatedOn(createdOn time.Time) ModificationBuilder {
	app.pCreatedOn = &createdOn
	return app
}

// Now builds a new Modification instance
func (app *modificationBuilder) Now() (Modification, error) {
	if app.pCreatedOn == nil {
		return nil, errors.New("the creation time is mandatory in order to build a Modification instance")
	}

	if app.name != "" && app.sig != nil && app.enc != nil {
		content := createContentWithNameAndSignatureAndEncryption(app.name, app.sig, app.enc)
		return createModification(content, *app.pCreatedOn), nil
	}

	if app.name != "" && app.sig != nil {
		content := createContentWithNameAndSignature(app.name, app.sig)
		return createModification(content, *app.pCreatedOn), nil
	}

	if app.name != "" && app.enc != nil {
		content := createContentWithNameAndEncryption(app.name, app.enc)
		return createModification(content, *app.pCreatedOn), nil
	}

	if app.sig != nil && app.enc != nil {
		content := createContentWithSignatureAndEncryption(app.sig, app.enc)
		return createModification(content, *app.pCreatedOn), nil
	}

	if app.name != "" {
		content := createContentWithName(app.name)
		return createModification(content, *app.pCreatedOn), nil
	}

	if app.sig != nil {
		content := createContentWithSignature(app.sig)
		return createModification(content, *app.pCreatedOn), nil
	}

	if app.enc != nil {
		content := createContentWithEncryption(app.enc)
		return createModification(content, *app.pCreatedOn), nil
	}

	return nil, errors.New("the Modification is invalid")
}
