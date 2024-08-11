package roots

import "github.com/steve-care-software/webx/engine/cursors/domain/loaders/identities/switchers/singles/namespaces/switchers/singles/blockchains/roots/units"

// Builder represents a root builder
type Builder interface {
	Create() Builder
	WithUnits(units units.Unit) Builder
	WithBaseDifficulty(baseDifficulty uint64) Builder
	WithIncreasePerSize(incrPerSize uint64) Builder
	WithSizeBlock(sizeBlock uint64) Builder
	Now() (Root, error)
}

// Root represents the root block
type Root interface {
	Units() units.Unit
	BaseDifficulty() uint64
	IncreasePerSize() uint64
	SizeBlock() uint64
}
