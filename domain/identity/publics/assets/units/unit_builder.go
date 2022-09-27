package units

import (
	"errors"
	"fmt"

	"github.com/steve-care-software/syntax/domain/identity/cryptography/hash"
	"github.com/steve-care-software/syntax/domain/identity/cryptography/signatures"
)

type unitBuilder struct {
	hashAdapter      hash.Adapter
	signatureAdapter signatures.RingSignatureAdapter
	content          Content
	signatures       []signatures.RingSignature
}

func createUnitBuilder(
	hashAdapter hash.Adapter,
	signatureAdapter signatures.RingSignatureAdapter,
) UnitBuilder {
	out := unitBuilder{
		hashAdapter:      hashAdapter,
		signatureAdapter: signatureAdapter,
		content:          nil,
		signatures:       nil,
	}

	return &out
}

// Create initializes the builder
func (app *unitBuilder) Create() UnitBuilder {
	return createUnitBuilder(
		app.hashAdapter,
		app.signatureAdapter,
	)
}

// WithContent adds content to the builder
func (app *unitBuilder) WithContent(content Content) UnitBuilder {
	app.content = content
	return app
}

// WithSignatures add signatures to the builder
func (app *unitBuilder) WithSignatures(signatures []signatures.RingSignature) UnitBuilder {
	app.signatures = signatures
	return app
}

func (app *unitBuilder) validateSignatures() error {
	msg := app.content.Hash().String()
	previous := app.content.Previous()
	if previous.IsGenesis() {
		if len(app.signatures) != 1 {
			str := fmt.Sprintf("there must be 1 signature in order to unlock the Genesis, %d provided", len(app.signatures))
			return errors.New(str)
		}

		genesis := previous.Genesis()
		prevOwner := genesis.Owner()
		isValid, err := app.signatureAdapter.ToVerification(app.signatures[0], msg, prevOwner)
		if err != nil {
			return err
		}

		if !isValid {
			str := fmt.Sprintf("the genesis (hash: %s) could not be unlocked with the provided signature", genesis.Hash().String())
			return errors.New(str)
		}

		return nil
	}

	units := previous.Units().List()
	if len(units) != len(app.signatures) {
		str := fmt.Sprintf("there must be the same amount of signatures (%d provided) as the amount of previous units (%d provided)", len(app.signatures), len(units))
		return errors.New(str)
	}

	for idx, oneUnit := range units {
		prevOwner := oneUnit.Content().Owner()
		isValid, err := app.signatureAdapter.ToVerification(app.signatures[idx], msg, prevOwner)
		if err != nil {
			return err
		}

		if !isValid {
			str := fmt.Sprintf("the unit (index: %d, hash: %s) could not be unlocked with the matching signature", idx, oneUnit.Hash().String())
			return errors.New(str)
		}
	}

	return nil
}

// Now builds a new Unit instance
func (app *unitBuilder) Now() (Unit, error) {
	if app.content == nil {
		return nil, errors.New("the content is mandatory in order to build a Unit instance")
	}

	if app.signatures != nil && len(app.signatures) <= 0 {
		app.signatures = nil
	}

	if app.signatures == nil {
		return nil, errors.New("there must be at least 1 Signature in order to build a Unit instance")
	}

	err := app.validateSignatures()
	if err != nil {
		return nil, err
	}

	data := [][]byte{
		app.content.Hash().Bytes(),
	}

	for _, oneSignature := range app.signatures {
		data = append(data, []byte(oneSignature.String()))
	}

	hash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	return createUnit(*hash, app.content, app.signatures), nil
}
