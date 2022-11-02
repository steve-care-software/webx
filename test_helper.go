package main

import "github.com/steve-care-software/webx/domain/cryptography/hash"

func valueToHashStringForTests(value string) string {
	pHash, err := hash.NewAdapter().FromBytes([]byte(value))
	if err != nil {
		panic(err)
	}

	return pHash.String()
}
