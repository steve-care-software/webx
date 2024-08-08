package roots

import "github.com/steve-care-software/webx/engine/cursors/domain/loaders/blockchains/roots/units"

// Root represents the root block
type Root interface {
	Units() units.Unit
	BaseDifficulty() uint64
	IncreasePerSize() uint64
	SizeBock() uint64
}
