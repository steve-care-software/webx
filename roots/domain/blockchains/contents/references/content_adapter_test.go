package references

import (
	"reflect"
	"testing"
)

func TestContentAdapter_withActive_withPending_withDeleted_Success(t *testing.T) {
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

	instance, err := NewContentBuilder().Create().WithActive(active).WithPendings(pendings).WithDeleted(deleted).Now()
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	adapter := NewContentAdapter()
	content, err := adapter.ToContent(instance)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retInstance, err := adapter.ToInstance(content)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !reflect.DeepEqual(instance, retInstance) {
		t.Errorf("the returned content instance is invalid")
		return
	}
}

func TestContentAdapter_withActive_withPending_Success(t *testing.T) {
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

	instance, err := NewContentBuilder().Create().WithActive(active).WithPendings(pendings).Now()
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	adapter := NewContentAdapter()
	content, err := adapter.ToContent(instance)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retInstance, err := adapter.ToInstance(content)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !reflect.DeepEqual(instance, retInstance) {
		t.Errorf("the returned content instance is invalid")
		return
	}
}

func TestContentAdapter_withActive_withDeleted_Success(t *testing.T) {
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

	deleted, err := NewContentKeysBuilder().Create().WithList([]ContentKey{
		NewContentKeyForTests(),
	}).Now()
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	instance, err := NewContentBuilder().Create().WithActive(active).WithDeleted(deleted).Now()
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	adapter := NewContentAdapter()
	content, err := adapter.ToContent(instance)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retInstance, err := adapter.ToInstance(content)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !reflect.DeepEqual(instance, retInstance) {
		t.Errorf("the returned content instance is invalid")
		return
	}
}

func TestContentAdapter_withPending_withDeleted_Success(t *testing.T) {
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

	instance, err := NewContentBuilder().Create().WithPendings(pendings).WithDeleted(deleted).Now()
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	adapter := NewContentAdapter()
	content, err := adapter.ToContent(instance)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retInstance, err := adapter.ToInstance(content)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !reflect.DeepEqual(instance, retInstance) {
		t.Errorf("the returned content instance is invalid")
		return
	}
}

func TestContentAdapter_Success(t *testing.T) {
	instance, err := NewContentFactory().Create()
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	adapter := NewContentAdapter()
	content, err := adapter.ToContent(instance)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retInstance, err := adapter.ToInstance(content)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !reflect.DeepEqual(instance, retInstance) {
		t.Errorf("the returned content instance is invalid")
		return
	}
}
