package grammars

import (
	"github.com/steve-care-software/webx/roots/domain/blockchains/cryptography/hash"
)

type grammar struct {
	hash     hash.Hash
	root     hash.Hash
	channels []hash.Hash
}

func createGrammar(
	hash hash.Hash,
	root hash.Hash,
) Grammar {
	return createGrammarInternally(hash, root, nil)
}

func createGrammarWithChannels(
	hash hash.Hash,
	root hash.Hash,
	channels []hash.Hash,
) Grammar {
	return createGrammarInternally(hash, root, channels)
}

func createGrammarInternally(
	hash hash.Hash,
	root hash.Hash,
	channels []hash.Hash,
) Grammar {
	out := grammar{
		hash:     hash,
		root:     root,
		channels: channels,
	}

	return &out
}

// Hash returns the hash
func (obj *grammar) Hash() hash.Hash {
	return obj.hash
}

// Root returns the root
func (obj *grammar) Root() hash.Hash {
	return obj.root
}

// HasChannels returns true if there is channels, false otherwise
func (obj *grammar) HasChannels() bool {
	return obj.channels != nil
}

// Channels returns the channels, if any
func (obj *grammar) Channels() []hash.Hash {
	return obj.channels
}
