package grammars

import "github.com/steve-care-software/webx/domain/cryptography/hash"

type grammar struct {
	hash     hash.Hash
	root     Token
	channels Channels
}

func createGrammar(
	hash hash.Hash,
	root Token,
) Grammar {
	return createGrammarInternally(hash, root, nil)
}

func createGrammarWithChannels(
	hash hash.Hash,
	root Token,
	channels Channels,
) Grammar {
	return createGrammarInternally(hash, root, channels)
}

func createGrammarInternally(
	hash hash.Hash,
	root Token,
	channels Channels,
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

// Root returns the root token
func (obj *grammar) Root() Token {
	return obj.root
}

// HasChannels returns true if there is channels, false otherwise
func (obj *grammar) HasChannels() bool {
	return obj.channels != nil
}

// Channels returns the channels, if any
func (obj *grammar) Channels() Channels {
	return obj.channels
}
