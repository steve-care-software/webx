package identities

import (
	"time"

	"github.com/steve-care-software/syntax/domain/syntax/databases"
	"github.com/steve-care-software/syntax/domain/syntax/databases/cryptography/encryptions/keys"
	"github.com/steve-care-software/syntax/domain/syntax/databases/cryptography/signatures"
	"github.com/steve-care-software/syntax/domain/syntax/identities/modifications"
)

type identity struct {
	name          string
	sig           signatures.PrivateKey
	enc           keys.PrivateKey
	createdOn     time.Time
	databases     databases.Databases
	modifications modifications.Modifications
}

func createIdentity(
	name string,
	sig signatures.PrivateKey,
	enc keys.PrivateKey,
	createdOn time.Time,
) Identity {
	return createIdentityInternally(name, sig, enc, createdOn, nil, nil)
}

func createIdentityWithDatabases(
	name string,
	sig signatures.PrivateKey,
	enc keys.PrivateKey,
	createdOn time.Time,
	databases databases.Databases,
) Identity {
	return createIdentityInternally(name, sig, enc, createdOn, databases, nil)
}

func createIdentityWithModifications(
	name string,
	sig signatures.PrivateKey,
	enc keys.PrivateKey,
	createdOn time.Time,
	modifications modifications.Modifications,
) Identity {
	return createIdentityInternally(name, sig, enc, createdOn, nil, modifications)
}

func createIdentityWithDatabasesAndModifications(
	name string,
	sig signatures.PrivateKey,
	enc keys.PrivateKey,
	createdOn time.Time,
	databases databases.Databases,
	modifications modifications.Modifications,
) Identity {
	return createIdentityInternally(name, sig, enc, createdOn, databases, modifications)
}

func createIdentityInternally(
	name string,
	sig signatures.PrivateKey,
	enc keys.PrivateKey,
	createdOn time.Time,
	databases databases.Databases,
	modifications modifications.Modifications,
) Identity {
	out := identity{
		name:          name,
		sig:           sig,
		enc:           enc,
		createdOn:     createdOn,
		databases:     databases,
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
	return obj.sig
}

// Encryption returns the encryption's pk
func (obj *identity) Encryption() keys.PrivateKey {
	return obj.enc
}

// CreatedOn returns the creation time
func (obj *identity) CreatedOn() time.Time {
	return obj.createdOn
}

// HasDatabases returns true if there is databases, false otherwise
func (obj *identity) HasDatabases() bool {
	return obj.databases != nil
}

// Databases returns the databases, if any
func (obj *identity) Databases() databases.Databases {
	return obj.databases
}

// HasModifications returns true if there is modifications, false otherwise
func (obj *identity) HasModifications() bool {
	return obj.modifications != nil
}

// Modifications returns the modifications, if any
func (obj *identity) Modifications() modifications.Modifications {
	return obj.modifications
}
