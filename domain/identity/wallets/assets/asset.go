package assets

import (
	uuid "github.com/satori/go.uuid"
	"github.com/steve-care-software/syntax/domain/identity/cryptography/signatures"
	"github.com/steve-care-software/syntax/domain/identity/units"
)

type asset struct {
	id   uuid.UUID
	pk   signatures.PrivateKey
	unit units.Unit
	ring []signatures.PublicKey
}

func createAsset(
	id uuid.UUID,
	pk signatures.PrivateKey,
	unit units.Unit,
	ring []signatures.PublicKey,
) Asset {
	out := asset{
		id:   id,
		pk:   pk,
		unit: unit,
		ring: ring,
	}

	return &out
}

// ID returns the id
func (obj *asset) ID() uuid.UUID {
	return obj.id
}

// PrivateKey returns the pk
func (obj *asset) PrivateKey() signatures.PrivateKey {
	return obj.pk
}

// Unit returns the unit
func (obj *asset) Unit() units.Unit {
	return obj.unit
}

// Ring returns the ring
func (obj *asset) Ring() []signatures.PublicKey {
	return obj.ring
}
