package encryptors

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
)

type encryptor struct {
	pk     rsa.PrivateKey
	pubKey PublicKey
}

func createEncryptor(
	pk rsa.PrivateKey,
	pubKey PublicKey,
) Encryptor {
	out := encryptor{
		pk:     pk,
		pubKey: pubKey,
	}

	return &out
}

// Decrypt decrypts a cipher
func (obj *encryptor) Decrypt(cipher []byte) ([]byte, error) {
	h := sha256.New()
	decrypted, err := rsa.DecryptOAEP(h, rand.Reader, &obj.pk, cipher, []byte(""))
	if err != nil {
		return nil, err
	}

	return decrypted, nil
}

// Public returns the public key
func (obj *encryptor) Public() PublicKey {
	return obj.pubKey
}

// Key returns the pk
func (obj *encryptor) Key() rsa.PrivateKey {
	return obj.pk
}
