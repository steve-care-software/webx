package references

import (
	"reflect"
	"testing"
)

func TestBlockchainAdapter_Success(t *testing.T) {
	blockchain := NewBlockchainForTests()
	adapter := NewBlockchainAdapter()
	content, err := adapter.ToContent(blockchain)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retBlockchain, err := adapter.ToBlockchain(content)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !reflect.DeepEqual(blockchain, retBlockchain) {
		t.Errorf("the returned blockchain is invalid")
		return
	}
}
