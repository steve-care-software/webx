package references

import (
	"reflect"
	"testing"
	"time"

	"github.com/steve-care-software/webx/blockchains/domain/cryptography/hash"
)

func TestBlockchainKeyAdapter_Success(t *testing.T) {
	pHash, err := hash.NewAdapter().FromBytes([]byte("this is some data"))
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	from := uint(233456)
	length := uint(345667899)
	pointer, err := NewPointerBuilder().Create().From(from).WithLength(length).Now()
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	createdOn := time.Now().UTC()
	blockchainKey, err := NewBlockchainKeyBuilder().Create().
		WithHash(*pHash).
		WithContent(pointer).
		CreatedOn(createdOn).
		Now()

	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	adapter := NewBlockchainKeyAdapter()
	content, err := adapter.ToContent(blockchainKey)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retBlockchainKey, err := adapter.ToBlockchainKey(content)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !reflect.DeepEqual(blockchainKey, retBlockchainKey) {
		t.Errorf("the returned blockchainKey is invalid")
		return
	}
}
