package accounts

import (
	"github.com/steve-care-software/identity/domain/accounts/encryptors"
	"github.com/steve-care-software/identity/domain/accounts/signers"
)

type account struct {
	username  string
	encryptor encryptors.Encryptor
	signer    signers.Signer
}

func createAccount(
	username string,
	encryptor encryptors.Encryptor,
	signer signers.Signer,
) Account {
	return createAccountInternally(username, encryptor, signer)
}

func createAccountInternally(
	username string,
	encryptor encryptors.Encryptor,
	signer signers.Signer,
) Account {
	out := account{
		username:  username,
		encryptor: encryptor,
		signer:    signer,
	}

	return &out
}

// Username returns the username
func (obj *account) Username() string {
	return obj.username
}

// Encryptor returns the encryptor
func (obj *account) Encryptor() encryptors.Encryptor {
	return obj.encryptor
}

// Signer returns the signer
func (obj *account) Signer() signers.Signer {
	return obj.signer
}
