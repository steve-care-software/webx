package references

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/steve-care-software/webx/blockchains/domain/cryptography/hash"
)

// NewBlockchainForTests creates a new blockchain for tests
func NewBlockchainForTests() Blockchain {
	chain := NewBlockchainKeyForTests()
	blocks, err := NewBlockchainKeysBuilder().Create().WithList([]BlockchainKey{
		NewBlockchainKeyForTests(),
		NewBlockchainKeyForTests(),
		NewBlockchainKeyForTests(),
		NewBlockchainKeyForTests(),
		NewBlockchainKeyForTests(),
	}).Now()
	if err != nil {
		panic(err)
	}

	trx, err := NewBlockchainKeysBuilder().Create().WithList([]BlockchainKey{
		NewBlockchainKeyForTests(),
		NewBlockchainKeyForTests(),
		NewBlockchainKeyForTests(),
		NewBlockchainKeyForTests(),
		NewBlockchainKeyForTests(),
	}).Now()
	if err != nil {
		panic(err)
	}

	ins, err := NewBlockchainBuilder().Create().WithChain(chain).WithBlocks(blocks).WithTransactions(trx).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewContentKeyForTests creates a new content key for tests
func NewContentKeyForTests() ContentKey {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	pHash, err := hash.NewAdapter().FromBytes([]byte(fmt.Sprintf("this is some data %d", r1.Int())))
	if err != nil {
		panic(err)
	}

	pTrxHash, err := hash.NewAdapter().FromBytes([]byte(fmt.Sprintf("this is some transaction data %d", r1.Int())))
	if err != nil {
		panic(err)
	}

	from := uint(r1.Intn(233456))
	length := uint(r1.Intn(22323)) + 1
	pointer, err := NewPointerBuilder().Create().From(from).WithLength(length).Now()
	if err != nil {
		panic(err)
	}

	createdOn := time.Now().UTC()
	ins, err := NewContentKeyBuilder().Create().
		WithHash(*pHash).
		WithKind(43).
		WithContent(pointer).
		WithTransaction(*pTrxHash).
		CreatedOn(createdOn).
		Now()

	if err != nil {
		panic(err)
	}

	return ins
}

// NewBlockchainKeyForTests creates a new blockchain key for tests
func NewBlockchainKeyForTests() BlockchainKey {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	pHash, err := hash.NewAdapter().FromBytes([]byte(fmt.Sprintf("this is some data %d", r1.Int())))
	if err != nil {
		panic(err)
	}

	from := uint(r1.Intn(233456))
	length := uint(r1.Intn(22323)) + 1
	pointer, err := NewPointerBuilder().Create().From(from).WithLength(length).Now()
	if err != nil {
		panic(err)
	}

	createdOn := time.Now().UTC()
	ins, err := NewBlockchainKeyBuilder().Create().
		WithHash(*pHash).
		WithContent(pointer).
		CreatedOn(createdOn).
		Now()

	if err != nil {
		panic(err)
	}

	return ins
}
