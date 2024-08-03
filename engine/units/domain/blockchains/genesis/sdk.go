package genesis

import (
	"github.com/steve-care-software/webx/engine/databases/hashes/domain/hash"
	"github.com/steve-care-software/webx/engine/units/domain/governances"
	"github.com/steve-care-software/webx/engine/units/domain/units/clears"
)

// Genesis represents the genesis
type Genesis interface {
	Hash() hash.Hash
	Unit() clears.Clear
	Governance() governances.Governance
	BaseDifficulty() uint8
	Increment() uint8
}
