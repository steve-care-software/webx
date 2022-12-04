package tokens

import "github.com/steve-care-software/webx/domain/cryptography/hash"

type token struct {
	hash   hash.Hash
	lines  Lines
	suites Suites
}

func createToken(
	hash hash.Hash,
	lines Lines,
) Token {
	return createTokenInternally(hash, lines, nil)
}

func createTokenWithSuites(
	hash hash.Hash,
	lines Lines,
	suites Suites,
) Token {
	return createTokenInternally(hash, lines, suites)
}

func createTokenInternally(
	hash hash.Hash,
	lines Lines,
	suites Suites,
) Token {
	out := token{
		hash:   hash,
		lines:  lines,
		suites: suites,
	}

	return &out
}

// Hash returns the hash
func (obj *token) Hash() hash.Hash {
	return obj.hash
}

// Lines returns the lines
func (obj *token) Lines() Lines {
	return obj.lines
}

// HasSuites returns true if there is suites, false otherwise
func (obj *token) HasSuites() bool {
	return obj.suites != nil
}

// Suites returns the suites, if any
func (obj *token) Suites() Suites {
	return obj.suites
}
