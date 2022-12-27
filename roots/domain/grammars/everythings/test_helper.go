package everythings

import "github.com/steve-care-software/webx/databases/domain/cryptography/hash"

// NewEverythingForTests creates a new everything for tests
func NewEverythingForTests() Everything {
	pHash, err := hash.NewAdapter().FromBytes([]byte("this is an hash"))
	if err != nil {
		panic(err)
	}

	pException, err := hash.NewAdapter().FromBytes([]byte("this is an exception hash"))
	if err != nil {
		panic(err)
	}

	ins, err := NewBuilder().Create().WithHash(*pHash).WithException(*pException).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewEverythingWithEscapeForTests creates a new everything with escape for tests
func NewEverythingWithEscapeForTests() Everything {
	pHash, err := hash.NewAdapter().FromBytes([]byte("this is an hash"))
	if err != nil {
		panic(err)
	}

	pException, err := hash.NewAdapter().FromBytes([]byte("this is an exception hash"))
	if err != nil {
		panic(err)
	}

	pEscape, err := hash.NewAdapter().FromBytes([]byte("this is an escape hash"))
	if err != nil {
		panic(err)
	}

	ins, err := NewBuilder().Create().WithHash(*pHash).WithException(*pException).WithEscape(*pEscape).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
