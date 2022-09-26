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
}

// AssetBuilder represents an asset builder
type AssetBuilder interface {
	Create() AssetBuilder
	WithID(id uuid.UUID) AssetBuilder
	WithPrivateKey(pk signatures.PrivateKey) AssetBuilder
	WithUnit(unit units.Unit) AssetBuilder
	Now() (Asset, error)
}

// Asset represents an asset
type Asset interface {
	ID() uuid.UUID
	PrivateKey() signatures.PrivateKey
	Unit() units.Unit
}
