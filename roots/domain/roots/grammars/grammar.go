package grammars

import (
	"github.com/steve-care-software/webx/roots/domain/blockchains/cryptography/hash"
	"github.com/steve-care-software/webx/roots/domain/blockchains/cryptography/hashtrees"
)

type grammar struct {
	hash    hash.Hash
	name    string
	history hashtrees.HashTree
}

func createGrammar(
	hash hash.Hash,
	name string,
) Grammar {
	return createGrammarInternally(hash, name, nil)
}

func createGrammarWithHistory(
	hash hash.Hash,
	name string,
	history hashtrees.HashTree,
) Grammar {
	return createGrammarInternally(hash, name, history)
}

func createGrammarInternally(
	hash hash.Hash,
	name string,
	history hashtrees.HashTree,
) Grammar {
	out := grammar{
		hash:    hash,
		name:    name,
		history: history,
	}

	return &out
}

// Hash returns the hash
func (obj *grammar) Hash() hash.Hash {
	return obj.hash
}

// Name returns the name
func (obj *grammar) Name() string {
	return obj.name
}

// HasHistory returns true if there is an history, false otherwise
func (obj *grammar) HasHistory() bool {
	return obj.history != nil
}

// History returns the history, if any
func (obj *grammar) History() hashtrees.HashTree {
	return obj.history
}
