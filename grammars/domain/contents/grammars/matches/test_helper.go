package matches

import (
	"fmt"

	"github.com/steve-care-software/webx/blockchains/domain/cryptography/hash"
)

// NewMatchForTests createsa new match for tests
func NewMatchForTests(suitesAmount uint) Match {
	hashAdapter := hash.NewAdapter()
	pHash, err := hashAdapter.FromBytes([]byte("this is an hash"))
	if err != nil {
		panic(err)
	}

	pTokenHash, err := hashAdapter.FromBytes([]byte("this is an hash"))
	if err != nil {
		panic(err)
	}

	suites := []hash.Hash{}
	for i := 0; i < int(suitesAmount); i++ {
		pSuiteHash, err := hashAdapter.FromBytes([]byte(fmt.Sprintf("this is a suite hash: %d", i)))
		if err != nil {
			panic(err)
		}

		suites = append(suites, *pSuiteHash)
	}

	ins, err := NewBuilder().Create().WithHash(*pHash).WithToken(*pTokenHash).WithSuites(suites).Now()
	if err != nil {
		panic(err)
	}

	return ins

}
