package layers

import (
	"reflect"
	"testing"
)

func TestIdentity_withSigner_Success(t *testing.T) {
	signer := NewSignerWithSignForTests("mySign")
	identity := NewIdentityWithSignerForTests(signer)

	if !identity.IsSigner() {
		t.Errorf("the identity was expected to contain a signer")
		return
	}

	if identity.IsEncryptor() {
		t.Errorf("the identity was expected to NOT contain an encryptor")
		return
	}

	retSigner := identity.Signer()
	if !reflect.DeepEqual(signer, retSigner) {
		t.Errorf("the returned signer is invalid")
		return
	}
}

func TestIdentity_withEncryptor_Success(t *testing.T) {
	variable := "myVariable"
	encryptor := NewEncryptorWithEncryptForTests(variable)
	identity := NewIdentityWithEncryptorForTests(encryptor)

	if identity.IsSigner() {
		t.Errorf("the identity was expected to NOT contain a signer")
		return
	}

	if !identity.IsEncryptor() {
		t.Errorf("the identity was expected to contain an encryptor")
		return
	}

	retEncryptor := identity.Encryptor()
	if !reflect.DeepEqual(encryptor, retEncryptor) {
		t.Errorf("the returned encryptor is invalid")
		return
	}
}

func TestIdentity_withParam_returnsError(t *testing.T) {
	_, err := NewIdentityBuilder().Create().Now()
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}
