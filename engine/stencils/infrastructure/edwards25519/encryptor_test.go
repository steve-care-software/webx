package edwards25519

import (
	"bytes"
	"testing"
)

func TestEncryptor_Success(t *testing.T) {
	password := []byte("this is a password")
	message := []byte("this is a message")
	encryptor := NewEncryptor()
	cipher, err := encryptor.Encrypt(message, password)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned")
		return
	}

	retMessage, err := encryptor.Decrypt(cipher, password)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned")
		return
	}

	if !bytes.Equal(message, retMessage) {
		t.Errorf("the returned message is invalid")
		return
	}
}

func TestEncryptor__decryptsWithInvalidPassword_ReturnsError(t *testing.T) {
	password := []byte("this is a password")
	message := []byte("this is a message")
	encryptor := NewEncryptor()
	cipher, err := encryptor.Encrypt(message, password)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned")
		return
	}

	invalidPassword := []byte("invalid password")
	retMessage, err := encryptor.Decrypt(cipher, invalidPassword)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned")
		return
	}

	if bytes.Equal(message, retMessage) {
		t.Errorf("the returned message was expected to be valid, but it was valid")
		return
	}
}
