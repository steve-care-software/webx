package encryptors

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
)

type publicKey struct {
	key rsa.PublicKey
}

func createPublicKey(key rsa.PublicKey) PublicKey {
	out := publicKey{
		key: key,
	}

	return &out
}

// Encrypt encrypts a message using the public key
func (obj *publicKey) Encrypt(msg []byte) ([]byte, error) {
	h := sha256.New()
	return rsa.EncryptOAEP(h, rand.Reader, &obj.key, msg, []byte(""))
}

// Key returns the public key
func (obj *publicKey) Key() rsa.PublicKey {
	return obj.key
}
