package grammars

import "github.com/steve-care-software/webx/blockchains/domain/cryptography/hash"

type token struct {
	hash   hash.Hash
	name   string
	block  Block
	suites Suites
}

func createToken(
	hash hash.Hash,
	name string,
	block Block,
) Token {
	return createTokenInternally(hash, name, block, nil)
}

func createTokenWithSuites(
	hash hash.Hash,
	name string,
	block Block,
	suites Suites,
) Token {
	return createTokenInternally(hash, name, block, suites)
}

func createTokenInternally(
	hash hash.Hash,
	name string,
	block Block,
	suites Suites,
) Token {
	out := token{
		hash:   hash,
		name:   name,
		block:  block,
		suites: suites,
	}

	return &out
}

// Hash returns the hash
func (obj *token) Hash() hash.Hash {
	return obj.hash
}

// Name returns the name
func (obj *token) Name() string {
	return obj.name
}

// Block returns the block
func (obj *token) Block() Block {
	return obj.block
}

// HasSuites returns true if there is suites, false otherwise
func (obj *token) HasSuites() bool {
	return obj.suites != nil
}

// Suites returns the suites, if any
func (obj *token) Suites() Suites {
	return obj.suites
}
