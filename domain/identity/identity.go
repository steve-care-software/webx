package identity

import (
	uuid "github.com/satori/go.uuid"
	"github.com/steve-care-software/syntax/domain/identity/cryptography/signatures"
	"github.com/steve-care-software/syntax/domain/identity/publics"
	"github.com/steve-care-software/syntax/domain/identity/wallets"
)

type identity struct {
	id      uuid.UUID
	public  publics.Public
	pk      signatures.PrivateKey
	wallets wallets.Wallets
}

func createIdentity(
	id uuid.UUID,
	public publics.Public,
	pk signatures.PrivateKey,
) Identity {
	return createIdentityInternally(id, public, pk, nil)
}

func createIdentityWithWallets(
	id uuid.UUID,
	public publics.Public,
	pk signatures.PrivateKey,
	wallets wallets.Wallets,
) Identity {
	return createIdentityInternally(id, public, pk, wallets)
}

func createIdentityInternally(
	id uuid.UUID,
	public publics.Public,
	pk signatures.PrivateKey,
	wallets wallets.Wallets,
) Identity {
	out := identity{
		id:      id,
		public:  public,
		pk:      pk,
		wallets: wallets,
	}

	return &out
}

// ID returns the id
func (obj *identity) ID() uuid.UUID {
	return obj.id
}

// Public returns the public
func (obj *identity) Public() publics.Public {
	return obj.public
}

// PrivateKey returns the signature's pk
func (obj *identity) PrivateKey() signatures.PrivateKey {
	return obj.pk
}

// HasWallets returns true if there is wallets, false otherwise
func (obj *identity) HasWallets() bool {
	return obj.wallets != nil
}

// Wallets returns the wallets, if any
func (obj *identity) Wallets() wallets.Wallets {
	return obj.wallets
}
