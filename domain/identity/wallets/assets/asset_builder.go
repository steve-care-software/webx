package assets

import (
	"bytes"
	"errors"
	"fmt"

	uuid "github.com/satori/go.uuid"
	"github.com/steve-care-software/syntax/domain/identity/cryptography/hash"
	"github.com/steve-care-software/syntax/domain/identity/cryptography/signatures"
	public_assets "github.com/steve-care-software/syntax/domain/identity/publics/assets"
)

type assetBuilder struct {
	hashAdapter hash.Adapter
	pID         *uuid.UUID
	pk          signatures.PrivateKey
	public      public_assets.Asset
	ring        []signatures.PublicKey
}

func createAssetBuilder(
	hashAdapter hash.Adapter,
) AssetBuilder {
	out := assetBuilder{
		hashAdapter: hashAdapter,
		pID:         nil,
		pk:          nil,
		public:      nil,
		ring:        nil,
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

// WithPublic adds a public asset to the builder
func (app *assetBuilder) WithPublic(public public_assets.Asset) AssetBuilder {
	app.public = public
	return app
}

// WithRing adds a ring to the builder
func (app *assetBuilder) WithRing(ring []signatures.PublicKey) AssetBuilder {
	app.ring = ring
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

	if app.public == nil {
		return nil, errors.New("the public Asset is mandatory in order to build an Asset instance")
	}

	if app.ring != nil && len(app.ring) <= 0 {
		app.ring = nil
	}

	if app.ring == nil {
		return nil, errors.New("there must be at least 1 PublicKey in the ring in order to build an Asset instance")
	}

	pubKey := app.pk.PublicKey()
	pubKeyHash, err := app.hashAdapter.FromBytes([]byte(pubKey.String()))
	if err != nil {
		return nil, err
	}

	ring := app.public.Unit().Content().Owner()
	if len(ring) != len(app.ring) {
		str := fmt.Sprintf("the unit contains %d hash in its owner's hashes, but the asset only contain %d PublicKey in its ring.  Those number should match", len(ring), len(app.ring))
		return nil, errors.New(str)
	}

	isUnitValid := false
	for _, oneHash := range ring {
		if bytes.Compare(oneHash.Bytes(), pubKeyHash.Bytes()) == 0 {
			isUnitValid = true
			break
		}
	}

	if !isUnitValid {
		return nil, errors.New("the provided PrivateKey does not have a PublicKey that is present in the Unit's owner hashes and therefore the provided PrivateKey cannot unlock the given Unit")
	}

	isRingValid := false
	for _, onePublicKey := range app.ring {
		if onePublicKey.Equals(pubKey) {
			isRingValid = true
			break
		}
	}

	if !isRingValid {
		return nil, errors.New("the provided PrivateKey does not have a PublicKey that is present in the Asset's PublicKey ring and therefore the provided PrivateKey cannot unlock the given Unit")
	}

	return createAsset(*app.pID, app.pk, app.public, app.ring), nil
}
