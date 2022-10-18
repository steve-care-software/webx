package identities

import (
	"time"

	"github.com/steve-care-software/syntax/domain/syntax/databases/cryptography/encryptions/keys"
	"github.com/steve-care-software/syntax/domain/syntax/databases/cryptography/signatures"
	"github.com/steve-care-software/syntax/domain/syntax/identities/modifications"
)

type identity struct {
	name          string
	sigPK         signatures.PrivateKey
	encPK         keys.PrivateKey
	modifications modifications.Modifications
}

func createIdentity(
	name string,
	sigPK signatures.PrivateKey,
	encPK keys.PrivateKey,
	modifications modifications.Modifications,
) Identity {
	return createIdentityInternally(name, sigPK, encPK, modifications)
}

func createIdentityInternally(
	name string,
	sigPK signatures.PrivateKey,
	encPK keys.PrivateKey,
	modifications modifications.Modifications,
) Identity {
	out := identity{
		name:          name,
		sigPK:         sigPK,
		encPK:         encPK,
		modifications: modifications,
	}

	return &out
}

// Name returns the name
func (obj *identity) Name() string {
	return obj.name
}

// Signature returns the signature's pk
func (obj *identity) Signature() signatures.PrivateKey {
	return obj.sigPK
}

// Encryption returns the encryption's pk
func (obj *identity) Encryption() keys.PrivateKey {
	return obj.encPK
}

// CreatedOn returns the creation time
func (obj *identity) CreatedOn() time.Time {
	return obj.modifications.First().CreatedOn()
}

// Modifications returns the modifications, if any
func (obj *identity) Modifications() modifications.Modifications {
	return obj.modifications
}
