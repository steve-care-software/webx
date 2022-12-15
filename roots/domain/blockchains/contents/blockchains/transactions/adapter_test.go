package transactions

import (
	"reflect"
	"testing"
)

func TestAdapter_Success(t *testing.T) {
	trx := NewTransactionForTests()
	adapter := NewAdapter()
	content, err := adapter.ToContent(trx)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retTrx, err := adapter.ToTransaction(content)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !reflect.DeepEqual(trx, retTrx) {
		t.Errorf("the returned transaction is invalid")
		return
	}
}
