package layers

import (
	"reflect"
	"testing"
)

func TestEncryptor_withDecrypt_Success(t *testing.T) {
	variable := "myVariable"
	encryptor := NewEncryptorWithDecryptForTests(variable)

	if !encryptor.IsDecrypt() {
		t.Errorf("the encryptor was expected to contain a decrypt")
		return
	}

	if encryptor.IsEncrypt() {
		t.Errorf("the encryptor was expected to NOT contain an encrypt")
		return
	}

	if encryptor.IsPublicKey() {
		t.Errorf("the encryptor was expected to NOT contain a publicKey")
		return
	}

	retVariable := encryptor.Decrypt()
	if !reflect.DeepEqual(variable, retVariable) {
		t.Errorf("the returned decrypt is invalid")
		return
	}
}

func TestEncryptor_withEncrypt_Success(t *testing.T) {
	variable := "myVariable"
	encryptor := NewEncryptorWithEncryptForTests(variable)

	if encryptor.IsDecrypt() {
		t.Errorf("the encryptor was expected to NOT contain a decrypt")
		return
	}

	if !encryptor.IsEncrypt() {
		t.Errorf("the encryptor was expected to contain an encrypt")
		return
	}

	if encryptor.IsPublicKey() {
		t.Errorf("the encryptor was expected to NOT contain a publicKey")
		return
	}

	retVariable := encryptor.Encrypt()
	if !reflect.DeepEqual(variable, retVariable) {
		t.Errorf("the returned encrypt is invalid")
		return
	}
}

func TestEncryptor_isPublicKey_Success(t *testing.T) {
	encryptor := NewEncryptorWithPublicKeyForTests()

	if encryptor.IsDecrypt() {
		t.Errorf("the encryptor was expected to NOT contain a decrypt")
		return
	}

	if encryptor.IsEncrypt() {
		t.Errorf("the encryptor was expected to NOT contain an encrypt")
		return
	}

	if !encryptor.IsPublicKey() {
		t.Errorf("the encryptor was expected to contain a publicKey")
		return
	}
}

func TestEncryptor_withoutParam_returnsError(t *testing.T) {
	_, err := NewEncryptorBuilder().Create().Now()
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}
