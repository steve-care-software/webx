package blockchains

import "github.com/steve-care-software/webx/roots/domain/blockchains/cryptography/hash"

// NewBlockchainForTests creates a new blockchain for tests
func NewBlockchainForTests() Blockchain {
	hashAdapter := hash.NewAdapter()
	pReference, err := hashAdapter.FromBytes([]byte("this is the reference hash"))
	if err != nil {
		panic(err)
	}

	pHead, err := hashAdapter.FromBytes([]byte("this is the head hash"))
	if err != nil {
		panic(err)
	}

	ins, err := NewBuilder().Create().WithReference(*pReference).WithHead(*pHead).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
