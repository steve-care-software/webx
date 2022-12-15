package transactions

import (
	"fmt"
	"math/big"
	"math/rand"
	"time"

	"github.com/steve-care-software/webx/roots/domain/blockchains/cryptography/hash"
)

// NewTransactionForTests creates a new transaction for tests
func NewTransactionForTests() Transaction {
	hashAdapter := hash.NewAdapter()
	pHash, err := hashAdapter.FromBytes([]byte("this is the hash"))
	if err != nil {
		panic(err)
	}

	pAsset, err := hashAdapter.FromBytes([]byte("this is the asset hash"))
	if err != nil {
		panic(err)
	}

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	pProof, _ := big.NewInt(0).SetString(fmt.Sprintf("%d%d%d%d%d%d", r1.Int(), r1.Int(), r1.Int(), r1.Int(), r1.Int(), r1.Int()), 10)

	ins, err := NewBuilder().Create().WithHash(*pHash).WithAsset(*pAsset).WithProof(*pProof).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
