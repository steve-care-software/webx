package keys

import (
	"time"

	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/identities/switchers/singles/keys/encryptors"
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/identities/switchers/singles/keys/signers"
)

// NewKeyForTests creates a new key for tests
func NewKeyForTests(
	encryptor encryptors.Encryptor,
	signer signers.Signer,
	createdOn time.Time,
) Key {
	ins, err := NewBuilder().Create().
		WithEncryptor(encryptor).
		WithSigner(signer).
		CreatedOn(createdOn).
		Now()

	if err != nil {
		panic(err)
	}

	return ins
}
