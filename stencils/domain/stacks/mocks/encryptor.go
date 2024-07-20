package mocks

import (
	"errors"

	"github.com/steve-care-software/datastencil/stencils/domain/encryptors"
)

type encryptor struct {
	encrypt map[string]map[string][]byte
	decrypt map[string]map[string][]byte
}

func createEncryptor(
	encrypt map[string]map[string][]byte,
	decrypt map[string]map[string][]byte,
) encryptors.Encryptor {
	out := encryptor{
		encrypt: encrypt,
		decrypt: decrypt,
	}

	return &out
}

// Encrypt encrypts a message
func (app *encryptor) Encrypt(message []byte, password []byte) ([]byte, error) {
	if mp, ok := app.encrypt[string(message)]; ok {
		if value, ok := mp[string(password)]; ok {
			return value, nil
		}

		return nil, errors.New("the password is invalid")
	}

	return nil, errors.New("the message is invalid")
}

// Decrypt decrypts a message
func (app *encryptor) Decrypt(cipher []byte, password []byte) ([]byte, error) {
	if mp, ok := app.decrypt[string(cipher)]; ok {
		if value, ok := mp[string(password)]; ok {
			return value, nil
		}

		return nil, errors.New("the password is invalid")
	}

	return nil, errors.New("the cipher is invalid")
}
