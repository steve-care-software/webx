package assets

import (
	"bytes"
	"errors"

	uuid "github.com/satori/go.uuid"
	"github.com/steve-care-software/syntax/domain/identity/cryptography/hash"
	"github.com/steve-care-software/syntax/domain/identity/cryptography/signatures"
	"github.com/steve-care-software/syntax/domain/identity/units"
)

type assetBuilder struct {
	hashAdapter hash.Adapter
	pID         *uuid.UUID
	pk          signatures.PrivateKey
	unit        units.Unit
}

func createAssetBuilder(
	hashAdapter hash.Adapter,
) AssetBuilder {
	out := assetBuilder{
		hashAdapter: hashAdapter,
		pID:         nil,
		pk:          nil,
		unit:        nil,
	}

	return &out
}

// Create initializes the builder
func (app *assetBuilder) Create() AssetBuilder {
	return createAssetBuilder(
		app.hashAdapter,
	)
}

// WithID adds an ID to the builder
func (app *assetBuilder) WithID(id uuid.UUID) AssetBuilder {
	app.pID = &id
	return app
}

// WithPrivateKey adds a pk to the builder
func (app *assetBuilder) WithPrivateKey(pk signatures.PrivateKey) AssetBuilder {
	app.pk = pk
	return app
}

// WithUnit adds a unit to the builder
func (app *assetBuilder) WithUnit(unit units.Unit) AssetBuilder {
	app.unit = unit
	return app
}

// Now builds a new Asset instance
func (app *assetBuilder) Now() (Asset, error) {
	if app.pID == nil {
		return nil, errors.New("the ID is mandatory in order to build an Asset instance")
	}

	if app.pk == nil {
		return nil, errors.New("the PrivateKey is mandatory in order to build an Asset instance")
	}

	if app.unit == nil {
		return nil, errors.New("the Unit is mandatory in order to build an Asset instance")
	}

	pubKey := app.pk.PublicKey()
	pubKeyHash, err := app.hashAdapter.FromBytes([]byte(pubKey.String()))
	if err != nil {
		return nil, err
	}

	isValid := false
	ring := app.unit.Content().Owner()
	for _, oneHash := range ring {
		if bytes.Compare(oneHash.Bytes(), pubKeyHash.Bytes()) == 0 {
			isValid = true
			break
		}
	}

	if !isValid {
		return nil, errors.New("the provided PrivateKey does not have a PublicKey that is present in the Unit's owner hashes and therefore the provided PrivateKey cannot unlock the given Unit")
	}

	return createAsset(*app.pID, app.pk, app.unit), nil
}
