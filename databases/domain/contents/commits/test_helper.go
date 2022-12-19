package commits

import (
	"github.com/steve-care-software/webx/databases/domain/cryptography/hash"
	"github.com/steve-care-software/webx/databases/domain/cryptography/hashtrees"
)

// NewCommitForTests creates a new commit for tests
func NewCommitForTests() Commit {
	pHash, err := hash.NewAdapter().FromBytes([]byte("this is an hash"))
	if err != nil {
		panic(err)
	}

	values, err := hashtrees.NewBuilder().Create().WithBlocks([][]byte{
		[]byte("first"),
		[]byte("second"),
		[]byte("third"),
	}).Now()
	if err != nil {
		panic(err)
	}

	ins, err := NewBuilder().Create().WithHash(*pHash).WithValues(values).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewCommitWithParentForTests creates a new commit with parent for tests
func NewCommitWithParentForTests() Commit {
	hashAdapter := hash.NewAdapter()
	pHash, err := hashAdapter.FromBytes([]byte("this is an hash"))
	if err != nil {
		panic(err)
	}

	values, err := hashtrees.NewBuilder().Create().WithBlocks([][]byte{
		[]byte("first"),
		[]byte("second"),
		[]byte("third"),
	}).Now()
	if err != nil {
		panic(err)
	}

	pParentHash, err := hashAdapter.FromBytes([]byte("this is a parent hash"))
	if err != nil {
		panic(err)
	}

	ins, err := NewBuilder().Create().WithHash(*pHash).WithValues(values).WithParent(*pParentHash).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
