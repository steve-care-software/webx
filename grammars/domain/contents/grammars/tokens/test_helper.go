package tokens

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/steve-care-software/webx/blockchains/domain/cryptography/hash"
)

// NewTokenForTests creates a new token for tests
func NewTokenForTests() Token {
	pHash, err := hash.NewAdapter().FromBytes([]byte("this is a token hash"))
	if err != nil {
		panic(err)
	}

	lines := NewLinesForTests(100)
	ins, err := NewBuilder().Create().WithHash(*pHash).WithLines(lines).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewLinesForTests creates a new lines for tests
func NewLinesForTests(amount int) Lines {
	list := []Line{}
	for i := 0; i < amount; i++ {
		s1 := rand.NewSource(time.Now().UnixNano())
		r1 := rand.New(s1)

		list = append(list, NewLineForTests(r1.Intn(200)+1))
	}

	ins, err := NewLinesBuilder().Create().WithList(list).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewLineForTests creates a new line for tests
func NewLineForTests(amount int) Line {
	list := []hash.Hash{}
	hashAdapter := hash.NewAdapter()
	for i := 0; i < amount; i++ {
		pHash, err := hashAdapter.FromBytes([]byte(fmt.Sprintf("this is an element: %d", i)))
		if err != nil {
			panic(err)
		}

		list = append(list, *pHash)
	}

	ins, err := NewLineBuilder().Create().WithElements(list).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
