package signers

import (
	"bytes"
	"testing"
)

func TestSigner_Success(t *testing.T) {
	// variables:
	pk := NewFactory().Create()
	pkBytes, err := pk.Bytes()
	if err != nil {
		t.Errorf("the returned error was expected to be nil, error returned: %s", err.Error())
		return
	}

	bckPK, err := NewAdapter().ToSigner(pkBytes)
	if err != nil {
		t.Errorf("the returned error was expected to be nil, error returned: %s", err.Error())
		return
	}

	bckPKBytes, err := bckPK.Bytes()
	if err != nil {
		t.Errorf("the returned error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !bytes.Equal(pkBytes, bckPKBytes) {
		t.Errorf("the Signers were expected to be the same.  Expected: %s, Returned: %s", pkBytes, bckPKBytes)
		return
	}
}
