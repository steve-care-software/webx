package channels

import "github.com/steve-care-software/webx/blockchains/domain/cryptography/hash"

// NewChannelForTests creates a new channel for tests
func NewChannelForTests() Channel {
	pHash, err := hash.NewAdapter().FromBytes([]byte("this is an hash"))
	if err != nil {
		panic(err)
	}

	pToken, err := hash.NewAdapter().FromBytes([]byte("this is a token hash"))
	if err != nil {
		panic(err)
	}

	ins, err := NewBuilder().Create().WithHash(*pHash).WithToken(*pToken).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewChannelWithPreviousForTests creates a new channel with previous for tests
func NewChannelWithPreviousForTests() Channel {
	pHash, err := hash.NewAdapter().FromBytes([]byte("this is an hash"))
	if err != nil {
		panic(err)
	}

	pToken, err := hash.NewAdapter().FromBytes([]byte("this is a token hash"))
	if err != nil {
		panic(err)
	}

	pPrevious, err := hash.NewAdapter().FromBytes([]byte("this is a previous hash"))
	if err != nil {
		panic(err)
	}

	ins, err := NewBuilder().Create().WithHash(*pHash).WithToken(*pToken).WithPrevious(*pPrevious).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewChannelWithNextForTests creates a new channel with next for tests
func NewChannelWithNextForTests() Channel {
	pHash, err := hash.NewAdapter().FromBytes([]byte("this is an hash"))
	if err != nil {
		panic(err)
	}

	pToken, err := hash.NewAdapter().FromBytes([]byte("this is a token hash"))
	if err != nil {
		panic(err)
	}

	pNext, err := hash.NewAdapter().FromBytes([]byte("this is a next hash"))
	if err != nil {
		panic(err)
	}

	ins, err := NewBuilder().Create().WithHash(*pHash).WithToken(*pToken).WithPrevious(*pNext).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewChannelWithPreviousAndNextForTests creates a new channel with previous and next for tests
func NewChannelWithPreviousAndNextForTests() Channel {
	pHash, err := hash.NewAdapter().FromBytes([]byte("this is an hash"))
	if err != nil {
		panic(err)
	}

	pToken, err := hash.NewAdapter().FromBytes([]byte("this is a token hash"))
	if err != nil {
		panic(err)
	}

	pPrevious, err := hash.NewAdapter().FromBytes([]byte("this is a previous hash"))
	if err != nil {
		panic(err)
	}

	pNext, err := hash.NewAdapter().FromBytes([]byte("this is a next hash"))
	if err != nil {
		panic(err)
	}

	ins, err := NewBuilder().Create().WithHash(*pHash).WithToken(*pToken).WithPrevious(*pPrevious).WithNext(*pNext).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
