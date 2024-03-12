package accounts

import (
	"github.com/steve-care-software/datastencil/domain/accounts/encryptors"
	"github.com/steve-care-software/datastencil/domain/accounts/signers"
)

// NewAccountForTests creates a new account for tests
func NewAccountForTests(username string, encryptor encryptors.Encryptor) Account {
	signer := signers.NewFactory().Create()
	ins, err := NewBuilder().Create().
		WithUsername(username).
		WithEncryptor(encryptor).
		WithSigner(signer).
		Now()

	if err != nil {
		panic(err)
	}

	return ins
}
