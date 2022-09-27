package assets

import (
	uuid "github.com/satori/go.uuid"
	"github.com/steve-care-software/syntax/domain/identity/cryptography/signatures"
	public_assets "github.com/steve-care-software/syntax/domain/identity/publics/assets"
)

type asset struct {
	id     uuid.UUID
	pk     signatures.PrivateKey
	public public_assets.Asset
	ring   []signatures.PublicKey
}

func createAsset(
	id uuid.UUID,
	pk signatures.PrivateKey,
	public public_assets.Asset,
	ring []signatures.PublicKey,
) Asset {
	out := asset{
		id:     id,
		pk:     pk,
		public: public,
		ring:   ring,
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

// Public returns the public asset
func (obj *asset) Public() public_assets.Asset {
	return obj.public
}

// Ring returns the ring
func (obj *asset) Ring() []signatures.PublicKey {
	return obj.ring
}
