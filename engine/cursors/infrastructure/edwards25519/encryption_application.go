package edwards25519

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"io"

	"github.com/steve-care-software/webx/engine/cursors/applications/encryptions"
)

type encryptionApplication struct {
}

func createEncryptionApplication() encryptions.Application {
	out := encryptionApplication{}
	return &out
}

// Encrypt encrypts a message
func (app *encryptionApplication) Encrypt(message []byte, password []byte) ([]byte, error) {
	key := app.hashPass(password)
	block, blockErr := aes.NewCipher(key)
	if blockErr != nil {
		return nil, blockErr
	}

	cipherBytes := make([]byte, aes.BlockSize+len(message))
	iv := cipherBytes[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(cipherBytes[aes.BlockSize:], message)

	return cipherBytes, nil
}

// Decrypt decrypts a cipher
func (app *encryptionApplication) Decrypt(cipherBytes []byte, password []byte) ([]byte, error) {
	key := app.hashPass(password)
	block, blockErr := aes.NewCipher(key)
	if blockErr != nil {
		return nil, blockErr
	}

	if len(cipherBytes) < aes.BlockSize {
		return nil, errors.New("the encrypted text cannot be decoded using this password: ciphertext block size is too short")
	}

	//IV needs to be unique, but doesn't have to be secure.
	//It's common to put it at the beginning of the ciphertext.
	iv := cipherBytes[:aes.BlockSize]
	cipherBytes = cipherBytes[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(cipherBytes, cipherBytes)

	// returns the decoded message:
	return cipherBytes, nil
}

func (app *encryptionApplication) hashPass(password []byte) []byte {
	hasher := curve.Hash()
	hasher.Write([]byte(password))
	return hasher.Sum(nil)
}
