package references

import (
	"reflect"
	"testing"
)

func TestAdapter_Success(t *testing.T) {
	active, err := NewContentKeysBuilder().Create().WithList([]ContentKey{
		NewContentKeyForTests(),
		NewContentKeyForTests(),
		NewContentKeyForTests(),
		NewContentKeyForTests(),
		NewContentKeyForTests(),
	}).Now()
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	pendings, err := NewContentKeysBuilder().Create().WithList([]ContentKey{
		NewContentKeyForTests(),
		NewContentKeyForTests(),
		NewContentKeyForTests(),
		NewContentKeyForTests(),
	}).Now()
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	deleted, err := NewContentKeysBuilder().Create().WithList([]ContentKey{
		NewContentKeyForTests(),
	}).Now()
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	contentIns, err := NewContentBuilder().Create().WithActive(active).WithPendings(pendings).WithDeleted(deleted).Now()
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	blockchain := NewBlockchainForTests()
	reference, err := NewBuilder().Create().WithContent(contentIns).WithBlockchain(blockchain).Now()
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	adapter := NewAdapter()
	content, err := adapter.ToContent(reference)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retReference, err := adapter.ToReference(content)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !reflect.DeepEqual(reference, retReference) {
		t.Errorf("the returned reference is invalid")
		return
	}
}
