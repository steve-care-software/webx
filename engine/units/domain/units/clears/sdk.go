package clears

import (
	"github.com/steve-care-software/webx/engine/databases/hashes/domain/hash"
	hideen_units "github.com/steve-care-software/webx/engine/units/domain/units"
)

// Clears represents a list of clear units
type Clears interface {
	Hash() hash.Hash
	List() []Clear
}

// Clear represents a clear unit
type Clear interface {
	Hash() hash.Hash
	Unit() hideen_units.Unit
	Answer() []byte
}
