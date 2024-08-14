package keys

import (
	"time"

	"github.com/steve-care-software/webx/engine/domain/identities/keys/encryptors"
	"github.com/steve-care-software/webx/engine/domain/identities/keys/signers"
)

type key struct {
	encryptor encryptors.Encryptor
	signer    signers.Signer
	createdOn time.Time
}

func createKey(
	encryptor encryptors.Encryptor,
	signer signers.Signer,
	createdOn time.Time,
) Key {
	out := key{
		encryptor: encryptor,
		signer:    signer,
		createdOn: createdOn,
	}

	return &out
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
