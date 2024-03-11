package retrieves

import "github.com/steve-care-software/datastencil/domain/hash"

type retrieve struct {
	hash        hash.Hash
	password    string
	credentials string
}

func createRetrieve(
	hash hash.Hash,
	password string,
	credentials string,
) Retrieve {
	out := retrieve{
		hash:        hash,
		password:    password,
		credentials: credentials,
	}

	return &out
}

// Hash returns the hash
func (obj *retrieve) Hash() hash.Hash {
	return obj.hash
}

// Password returns the password
func (obj *retrieve) Password() string {
	return obj.password
}

// Credentials returns the credentials
func (obj *retrieve) Credentials() string {
	return obj.credentials
}
