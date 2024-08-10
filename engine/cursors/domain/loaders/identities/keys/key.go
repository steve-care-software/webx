package keys

import (
	"time"

	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/identities/keys/encryptors"
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/identities/keys/profiles"
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/identities/keys/signers"
)

type key struct {
	profile   profiles.Profile
	encryptor encryptors.Encryptor
	signer    signers.Signer
	createdOn time.Time
}

func createKey(
	profile profiles.Profile,
	encryptor encryptors.Encryptor,
	signer signers.Signer,
	createdOn time.Time,
) Key {
	out := key{
		profile:   profile,
		encryptor: encryptor,
		signer:    signer,
		createdOn: createdOn,
	}

	return &out
}

// Profile returns the profile
func (obj *key) Profile() profiles.Profile {
	return obj.profile
}

// Encryptor returns the encryptor
func (obj *key) Encryptor() encryptors.Encryptor {
	return obj.encryptor
}

// Signer returns the signer
func (obj *key) Signer() signers.Signer {
	return obj.signer
}

// CreatedOn returns the creation time
func (obj *key) CreatedOn() time.Time {
	return obj.createdOn
}
