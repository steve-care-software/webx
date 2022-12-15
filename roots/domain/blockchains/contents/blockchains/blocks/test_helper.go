package blocks

import (
	"fmt"
	"math/big"
	"math/rand"
	"time"

	"github.com/steve-care-software/webx/roots/domain/blockchains/cryptography/hash"
	"github.com/steve-care-software/webx/roots/domain/blockchains/cryptography/hashtrees"
)

// NewBlockForTests creates a new block for tests
func NewBlockForTests(hasPrevious bool) Block {
	hashAdapter := hash.NewAdapter()
	pHash, err := hashAdapter.FromBytes([]byte("this is the hash"))
	if err != nil {
		panic(err)
	}

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	height := uint(r1.Int())

	s1 = rand.NewSource(time.Now().UnixNano())
	r1 = rand.New(s1)
	pNextScore, _ := big.NewInt(0).SetString(fmt.Sprintf("%d%d%d%d%d%d", r1.Int(), r1.Int(), r1.Int(), r1.Int(), r1.Int(), r1.Int()), 10)

	s1 = rand.NewSource(time.Now().UnixNano())
	r1 = rand.New(s1)
	pPendingScore, _ := big.NewInt(0).SetString(fmt.Sprintf("%d%d%d%d%d%d", r1.Int(), r1.Int(), r1.Int(), r1.Int(), r1.Int(), r1.Int()), 10)

	trx, err := hashtrees.NewBuilder().Create().WithBlocks([][]byte{
		[]byte("this"),
		[]byte("is"),
		[]byte("some"),
		[]byte("data"),
	}).Now()

	if err != nil {
		panic(err)
	}

	builder := NewBuilder().Create().
		WithHash(*pHash).
		WithHeight(height).
		WithNextScore(*pNextScore).
		WithPendingScope(*pPendingScore).
		WithTransactions(trx)

	if hasPrevious {
		pPrevHash, err := hashAdapter.FromBytes([]byte("this is the previous hash"))
		if err != nil {
			panic(err)
		}

		builder.WithPrevious(*pPrevHash)
	}

	ins, err := builder.Now()
	if err != nil {
		panic(err)
	}

	return ins

}
