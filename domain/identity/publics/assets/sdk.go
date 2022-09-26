package assets

import "github.com/steve-care-software/syntax/domain/identity/units"

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
	Now() (Asset, error)
}

// Asset represents a public asset
type Asset interface {
	IsUnit() bool
	Unit() units.Unit
}
