package credentials

import "github.com/steve-care-software/datastencil/domain/hash"

type credentials struct {
	hash     hash.Hash
	username string
	password string
}

func createCredentials(
	hash hash.Hash,
	username string,
	password string,
) Credentials {
	out := credentials{
		hash:     hash,
		username: username,
		password: password,
	}

	return &out
}

// Hash returns the hash
func (obj *credentials) Hash() hash.Hash {
	return obj.hash
}

// Username returns the username
func (obj *credentials) Username() string {
	return obj.username
}

// Password returns the password
func (obj *credentials) Password() string {
	return obj.password
}
