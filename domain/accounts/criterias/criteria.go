package criterias

import (
	"github.com/steve-care-software/datastencil/domain/accounts/encryptors"
	"github.com/steve-care-software/datastencil/domain/accounts/signers"
)

type criteria struct {
	signer    signers.Signer
	encryptor encryptors.Encryptor
	username  string
	password  []byte
}

func createCriteria(
	signer signers.Signer,
	encryptor encryptors.Encryptor,
	username string,
	password []byte,
) Criteria {
	out := criteria{
		signer:    signer,
		encryptor: encryptor,
		username:  username,
		password:  password,
	}

	return &out
}

// HasSigner returns true if there is a signer, false otherwise
func (obj *criteria) HasSigner() bool {
	return obj.signer != nil
}

// Signer returns the signer, if any
func (obj *criteria) Signer() signers.Signer {
	return obj.signer
}

// HasEncryptor returns true if there is an encryptor, false otherwise
func (obj *criteria) HasEncryptor() bool {
	return obj.encryptor != nil
}

// Encryptor returns the encryptor, if any
func (obj *criteria) Encryptor() encryptors.Encryptor {
	return obj.encryptor
}

// HasUsername returns true if there is a username, false otherwise
func (obj *criteria) HasUsername() bool {
	return obj.username != ""
}

// Username returns the username, if any
func (obj *criteria) Username() string {
	return obj.username
}

// HasPassword returns true if there is a password, false otherwise
func (obj *criteria) HasPassword() bool {
	return obj.password != nil
}

// Password returns the password, if any
func (obj *criteria) Password() []byte {
	return obj.password
}
