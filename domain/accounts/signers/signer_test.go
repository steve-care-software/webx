package signers

import (
	"testing"
)

func TestSigner_Success(t *testing.T) {
	// variables:
	pk := NewFactory().Create()
	pkStr := pk.String()
	bckPK, err := NewAdapter().ToSigner(pkStr)
	if err != nil {
		t.Errorf("the returned error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if pk.String() != bckPK.String() {
		t.Errorf("the Signers were expected to be the same.  Expected: %s, Returned: %s", pk.String(), bckPK.String())
		return
	}
}
