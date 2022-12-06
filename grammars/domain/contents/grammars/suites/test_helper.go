package suites

import "github.com/steve-care-software/webx/blockchains/domain/cryptography/hash"

// NewSuiteForTests creates a new suite for tests
func NewSuiteForTests(isValid bool) Suite {
	pHash, err := hash.NewAdapter().FromBytes([]byte("this is a suite hash"))
	if err != nil {
		panic(err)
	}

	content := []byte("this is some content")
	builder := NewBuilder().Create().WithContent(content).WithHash(*pHash)
	if isValid {
		builder.IsValid()
	}

	ins, err := builder.Now()
	if err != nil {
		panic(err)
	}

	return ins
}
