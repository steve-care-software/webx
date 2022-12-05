package keys

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
)

type publicKey struct {
	ky rsa.PublicKey
}

func createPublicKey(ky rsa.PublicKey) PublicKey {
	out := publicKey{
		ky: ky,
	}

	return &out
}

// Key returns the public key
func (obj *publicKey) Key() rsa.PublicKey {
	return obj.ky
}

// Encrypt encrypts a message using the public key
func (obj *publicKey) Encrypt(msg []byte) ([]byte, error) {
	h := sha256.New()
	return rsa.EncryptOAEP(h, rand.Reader, &obj.ky, msg, []byte(""))
}
