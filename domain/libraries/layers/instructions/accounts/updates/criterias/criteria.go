package criterias

import "github.com/steve-care-software/datastencil/domain/hash"

type criteria struct {
	hash            hash.Hash
	changeSigner    bool
	changeEncryptor bool
	username        string
	password        string
}

func createCriteria(
	hash hash.Hash,
	changeSigner bool,
	changeEncryptor bool,
	username string,
	password string,
) Criteria {
	out := criteria{
		hash:            hash,
		changeSigner:    changeSigner,
		changeEncryptor: changeEncryptor,
		username:        username,
		password:        password,
	}

	return &out
}

// ChangeSigner returns true if change signer, false otherwise
func (obj *criteria) Hash() hash.Hash {
	return obj.hash
}

// ChangeSigner returns true if change signer, false otherwise
func (obj *criteria) ChangeSigner() bool {
	return obj.changeSigner
}

// ChangeEncryptor returns true if change signer, false otherwise
func (obj *criteria) ChangeEncryptor() bool {
	return obj.changeEncryptor
}

// HasUsername returns true if there is username, false otherwise
func (obj *criteria) HasUsername() bool {
	return obj.username != ""
}

// Username returns the username, if any
func (obj *criteria) Username() string {
	return obj.username
}

// HasPassword returns true if there is password, false otherwise
func (obj *criteria) HasPassword() bool {
	return obj.password != ""
}

// Password returns the password, if any
func (obj *criteria) Password() string {
	return obj.password
}
