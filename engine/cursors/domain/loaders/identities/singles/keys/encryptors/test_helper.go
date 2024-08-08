package encryptors

import (
	"crypto/rand"
	"crypto/rsa"
)

// NewEncryptorForTests creates a new encryptor for tests
func NewEncryptorForTests(bitrate int) Encryptor {
	pPrivateKey, err := rsa.GenerateKey(rand.Reader, bitrate)
	if err != nil {
		panic(err)
	}

	ins, err := NewBuilder().Create().WithPK(*pPrivateKey).WithBitRate(bitrate).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
