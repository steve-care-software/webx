package assets

import (
	"github.com/steve-care-software/syntax/domain/identity/publics/assets/claims"
	"github.com/steve-care-software/syntax/domain/identity/publics/assets/units"
)

// Builder represents an assets builder
type Builder interface {
	Create() Builder
	WithList(list []Asset) Builder
	Now() (Assets, error)
}

// Assets represents public assets
type Assets interface {
	List() []Asset
}

// AssetBuilder represents an asset builder
type AssetBuilder interface {
	Create() AssetBuilder
	WithUnit(unit units.Unit) AssetBuilder
	WithClaim(claim claims.Claim) AssetBuilder
	Now() (Asset, error)
}

// Asset represents a public asset
type Asset interface {
	IsUnit() bool
	Unit() units.Unit
	IsClaim() bool
	Claim() claims.Claim
}
