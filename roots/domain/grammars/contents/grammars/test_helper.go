package grammars

import (
	"fmt"

	"github.com/steve-care-software/webx/blockchains/domain/cryptography/hash"
)

// NewGrammarForTests creates a new grammar for tests
func NewGrammarForTests() Grammar {
	pHash, err := hash.NewAdapter().FromBytes([]byte("this is a grammar hash"))
	if err != nil {
		panic(err)
	}

	pRoot, err := hash.NewAdapter().FromBytes([]byte("this is a root token hash"))
	if err != nil {
		panic(err)
	}

	ins, err := NewBuilder().Create().WithHash(*pHash).WithRoot(*pRoot).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewGrammarWithChannelsForTests creates a new grammar with channels for tests
func NewGrammarWithChannelsForTests(chanAmount int) Grammar {
	pHash, err := hash.NewAdapter().FromBytes([]byte("this is a grammar hash"))
	if err != nil {
		panic(err)
	}

	pRoot, err := hash.NewAdapter().FromBytes([]byte("this is a root token hash"))
	if err != nil {
		panic(err)
	}

	channels := []hash.Hash{}
	for i := 0; i < chanAmount; i++ {
		pChan, err := hash.NewAdapter().FromBytes([]byte(fmt.Sprintf("this is a channel hash: %d", i)))
		if err != nil {
			panic(err)
		}

		channels = append(channels, *pChan)
	}

	ins, err := NewBuilder().Create().WithHash(*pHash).WithRoot(*pRoot).WithChannels(channels).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
