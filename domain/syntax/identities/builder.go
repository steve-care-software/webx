package identities

import (
	"errors"

	"github.com/steve-care-software/syntax/domain/syntax/databases/cryptography/encryptions/keys"
	"github.com/steve-care-software/syntax/domain/syntax/databases/cryptography/signatures"
	"github.com/steve-care-software/syntax/domain/syntax/identities/modifications"
)

type builder struct {
	modifications modifications.Modifications
}

func createBuilder() Builder {
	out := builder{
		modifications: nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithModifications add modifications to the builder
func (app *builder) WithModifications(modifications modifications.Modifications) Builder {
	app.modifications = modifications
	return app
}

// Now builds a new Identity instance
func (app *builder) Now() (Identity, error) {
	if app.modifications == nil {
		return nil, errors.New("the modifications is mandatory in order to build an Identity instance")
	}

	name := ""
	var sigPK signatures.PrivateKey
	var encPK keys.PrivateKey

	list := app.modifications.List()
	for _, oneModification := range list {
		content := oneModification.Content()
		if content.HasSignature() {
			sigPK = content.Signature()
		}

		if content.HasEncryption() {
			encPK = content.Encryption()
		}

		if content.HasName() {
			name = content.Name()
		}
	}

	if sigPK == nil {
		return nil, errors.New("there is no modification that contains a signature's PrivateKey")
	}

	if encPK == nil {
		return nil, errors.New("there is no modification that contains an encryption's PrivateKey")
	}

	if name == "" {
		return nil, errors.New("the is no modification that contains a name")
	}

	return createIdentity(name, sigPK, encPK, app.modifications), nil
}
