package matches

import "github.com/steve-care-software/webx/blockchains/domain/cryptography/hash"

type match struct {
	hash   hash.Hash
	token  hash.Hash
	suites []hash.Hash
}

func createMatch(
	hash hash.Hash,
	token hash.Hash,
	suites []hash.Hash,
) Match {
	out := match{
		hash:   hash,
		token:  token,
		suites: suites,
	}

	return &out
}

// Hash returns the hash
func (obj *match) Hash() hash.Hash {
	return obj.hash
}

// Token returns the token
func (obj *match) Token() hash.Hash {
	return obj.token
}

// Suites returns the suites
func (obj *match) Suites() []hash.Hash {
	return obj.suites
}
