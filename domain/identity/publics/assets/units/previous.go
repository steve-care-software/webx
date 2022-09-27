package units

import (
	"github.com/steve-care-software/syntax/domain/identity/cryptography/hash"
	"github.com/steve-care-software/syntax/domain/identity/publics/assets/units/genesis"
)

type previous struct {
	units   Units
	genesis genesis.Genesis
}

func createPreviousWithUnits(
	units Units,
) Previous {
	return createPreviousInternally(units, nil)
}

func createPreviousWithGenesis(
	genesis genesis.Genesis,
) Previous {
	return createPreviousInternally(nil, genesis)
}

func createPreviousInternally(
	units Units,
	genesis genesis.Genesis,
) Previous {
	out := previous{
		units:   units,
		genesis: genesis,
	}

	return &out
}

// Hash returns the hash
func (obj *previous) Hash() hash.Hash {
	if obj.IsUnits() {
		return obj.Units().Hash()
	}

	return obj.Genesis().Hash()
}

// Amount returns the amount
func (obj *previous) Amount() uint64 {
	if obj.IsUnits() {
		return obj.Units().Amount()
	}

	return obj.Genesis().Supply()
}

// IsUnits returns true if there is units, false otherwise
func (obj *previous) IsUnits() bool {
	return obj.units != nil
}

// Units returns the units, if any
func (obj *previous) Units() Units {
	return obj.units
}

// IsGenesis returns true if there is genesis, false otherwise
func (obj *previous) IsGenesis() bool {
	return obj.genesis != nil
}

// Genesis returns the genesis, if any
func (obj *previous) Genesis() genesis.Genesis {
	return obj.genesis
}
