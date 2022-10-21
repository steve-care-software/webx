package keys

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
)

type privateKey struct {
	key    rsa.PrivateKey
	pubKey PublicKey
}

func createPrivateKey(key rsa.PrivateKey, pubKey PublicKey) PrivateKey {
	out := privateKey{
		key:    key,
		pubKey: pubKey,
	}

	return &out
}

// Key returns the key
func (obj *privateKey) Key() rsa.PrivateKey {
	return obj.key
}

// Public returns the public key
func (obj *privateKey) Public() PublicKey {
	return obj.pubKey
}

// Decrypt decrypts a cipher
func (obj *privateKey) Decrypt(cipher []byte) ([]byte, error) {
	h := sha256.New()
	decrypted, err := rsa.DecryptOAEP(h, rand.Reader, &obj.key, cipher, []byte(""))
	if err != nil {
		return nil, err
	}

	return decrypted, nil
}
