package uints

import "testing"

func TestAdapter_Success(t *testing.T) {
	value := uint64(3425324)
	adapter := NewAdapter()
	retAST, err := adapter.ToNFT(value)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	pRetValue, err := adapter.ToValue(retAST)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if value != *pRetValue {
		t.Errorf("the returned value was expected to be %d, %d returned", value, *pRetValue)
		return
	}
}
