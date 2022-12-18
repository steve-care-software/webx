package grammars

import (
	"github.com/steve-care-software/webx/databases/domain/cryptography/hash"
	"github.com/steve-care-software/webx/databases/domain/cryptography/hashtrees"
)

// NewGrammarForTests creates a new grammar for tests
func NewGrammarForTests() Grammar {
	pHash, err := hash.NewAdapter().FromBytes([]byte("this is an hash"))
	if err != nil {
		panic(err)
	}

	name := "this-is-a-name"
	ins, err := NewBuilder().Create().WithHash(*pHash).WithName(name).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewGrammarWithHistoryForTests creates a new grammar with history for tests
func NewGrammarWithHistoryForTests() Grammar {
	pHash, err := hash.NewAdapter().FromBytes([]byte("this is an hash"))
	if err != nil {
		panic(err)
	}

	name := "this-is-a-name"
	htIns, err := hashtrees.NewBuilder().WithBlocks([][]byte{
		[]byte("this"),
		[]byte("is"),
		[]byte("some"),
		[]byte("random"),
		[]byte("blocks"),
	}).Now()

	if err != nil {
		panic(err)
	}

	ins, err := NewBuilder().Create().WithHash(*pHash).WithName(name).WithHistory(htIns).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
