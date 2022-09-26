package assets

import (
	uuid "github.com/satori/go.uuid"
	"github.com/steve-care-software/syntax/domain/identity/cryptography/hash"
	"github.com/steve-care-software/syntax/domain/identity/cryptography/signatures"
	"github.com/steve-care-software/syntax/domain/identity/units"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// NewAssetBuilder creates a new asset builder
func NewAssetBuilder() AssetBuilder {
	hashAdapter := hash.NewAdapter()
	return createAssetBuilder(hashAdapter)
}

// Builder represents an assets builder
type Builder interface {
	Create() Builder
	WithList(list []Asset) Builder
	Now() (Assets, error)
}

// Assets represents assets
type Assets interface {
	List() []Asset
	FetchByUnits(unitHashes []hash.Hash) ([]Asset, error)
}

// AssetBuilder represents an asset builder
type AssetBuilder interface {
	Create() AssetBuilder
	WithID(id uuid.UUID) AssetBuilder
	WithUnit(unit units.Unit) AssetBuilder
	WithPrivateKey(pk signatures.PrivateKey) AssetBuilder
	WithRing(ring []signatures.PublicKey) AssetBuilder
	Now() (Asset, error)
}

// Asset represents an asset
type Asset interface {
	ID() uuid.UUID
	Unit() units.Unit
	PrivateKey() signatures.PrivateKey
	Ring() []signatures.PublicKey
}
