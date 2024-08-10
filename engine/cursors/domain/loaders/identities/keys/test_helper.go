package keys

import (
	"time"

	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/identities/keys/encryptors"
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/identities/keys/profiles"
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/identities/keys/signers"
)

// NewKeysForTests creates a new keys for tests
func NewKeysForTests(list []Key) Keys {
	ins, err := NewBuilder().Create().
		WithList(list).
		Now()

	if err != nil {
		panic(err)
	}

	return ins
}

// NewKeyForTests creates a new key for tests
func NewKeyForTests(
	profile profiles.Profile,
	encryptor encryptors.Encryptor,
	signer signers.Signer,
	createdOn time.Time,
) Key {
	ins, err := NewKeyBuilder().Create().
		WithProfile(profile).
		WithEncryptor(encryptor).
		WithSigner(signer).
		CreatedOn(createdOn).
		Now()

	if err != nil {
		panic(err)
	}

	return ins
}
