package signers

import (
	"bytes"
	"testing"
)

func TestSignature_Success(t *testing.T) {
	// variables:
	msg := []byte("this is a message to sign")
	pk := NewFactory().Create()

	// create the signature:
	sig, err := pk.Sign(msg)
	if err != nil {
		t.Errorf("the returned error was expected to be nil, error returned: %s", err.Error())
		return
	}

	// derive a PublicKey from the signature:
	derivedPubKey, err := sig.PublicKey(msg)
	if err != nil {
		t.Errorf("the returned error was expected to be nil, error returned: %s", err.Error())
		return
	}

	// make sure the original PublicKey and the derived PublicKey are the same:
	if !pk.PublicKey().Equals(derivedPubKey) {
		t.Errorf("the original PublicKey was expected to be the same as the derived PublicKey")
		return
	}

	// verify the signature:
	if !sig.Verify() {
		t.Errorf("the signature was expected to be verified using this message and PublicKey")
		return
	}

	// convert back and forth to string:
	sigBytes, err := sig.Bytes()
	if err != nil {
		t.Errorf("the returned error was expected to be nil, error returned: %s", err.Error())
		return
	}

	newSig, err := NewSignatureAdapter().ToSignature(sigBytes)
	if err != nil {
		t.Errorf("the returned error was expected to be nil, error returned: %s", err.Error())
		return
	}

	newSigBytes, err := newSig.Bytes()
	if err != nil {
		t.Errorf("the returned error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !bytes.Equal(sigBytes, newSigBytes) {
		t.Errorf("the signatures were expected to be the same.  Expected: %s, Actual: %s", sigBytes, newSigBytes)
		return
	}

}
