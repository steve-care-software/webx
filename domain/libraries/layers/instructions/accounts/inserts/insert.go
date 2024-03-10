package inserts

import "github.com/steve-care-software/datastencil/domain/hash"

type insert struct {
	hash     hash.Hash
	username string
	password string
}

func createInsert(
	hash hash.Hash,
	username string,
	password string,
) Insert {
	out := insert{
		hash:     hash,
		username: username,
		password: password,
	}

	return &out
}

// Hash returns the hash
func (obj *insert) Hash() hash.Hash {
	return obj.hash
}

// Username returns the username
func (obj *insert) Username() string {
	return obj.username
}

// Password returns the password
func (obj *insert) Password() string {
	return obj.password
}
