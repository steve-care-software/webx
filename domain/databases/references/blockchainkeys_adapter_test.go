package references

import (
	"reflect"
	"testing"
)

func TestBlockchainKeysAdapter_Success(t *testing.T) {
	list := []BlockchainKey{
		NewBlockchainKeyForTests(),
		NewBlockchainKeyForTests(),
		NewBlockchainKeyForTests(),
		NewBlockchainKeyForTests(),
		NewBlockchainKeyForTests(),
	}

	blockchainKeys, err := NewBlockchainKeysBuilder().Create().WithList(list).Now()
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	adapter := NewBlockchainKeysAdapter()
	content, err := adapter.ToContent(blockchainKeys)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retBlockchainKeys, err := adapter.ToBlockchainKeys(content)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !reflect.DeepEqual(blockchainKeys, retBlockchainKeys) {
		t.Errorf("the returned blockchainKeys is invalid")
		return
	}
}
