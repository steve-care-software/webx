package signers

import (
	"testing"
)

func TestPublicKey_Success(t *testing.T) {
	//variables:
	p := curve.Point().Base()

	// execute:
	pKey := createPublicKey(p)
	pubKeyBytes, err := pKey.Bytes()
	if err != nil {
		t.Errorf("the returned error was expected to be nil, error returned: %s", err.Error())
		return
	}

	samePubKey, err := NewPublicKeyAdapter().ToPublicKey(pubKeyBytes)
	if err != nil {
		t.Errorf("the returned error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !pKey.Equals(samePubKey) {
		t.Errorf("the public keys should be equal")
		return
	}
}
