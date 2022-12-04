package elements

import (
	"reflect"
	"testing"
)

func TestAdapter_withValue_Success(t *testing.T) {
	element := NewElementWithValueForTests()
	adapter := NewAdapter()
	content, err := adapter.ToContent(element)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retElement, err := adapter.ToElement(content)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !reflect.DeepEqual(element, retElement) {
		t.Errorf("the returned element is invalid")
		return
	}
}

func TestAdapter_withExternal_Success(t *testing.T) {
	element := NewElementWithExternalForTests()
	adapter := NewAdapter()
	content, err := adapter.ToContent(element)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retElement, err := adapter.ToElement(content)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !reflect.DeepEqual(element, retElement) {
		t.Errorf("the returned element is invalid")
		return
	}
}

func TestAdapter_withToken_Success(t *testing.T) {
	element := NewElementWithTokenForTests()
	adapter := NewAdapter()
	content, err := adapter.ToContent(element)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retElement, err := adapter.ToElement(content)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !reflect.DeepEqual(element, retElement) {
		t.Errorf("the returned element is invalid")
		return
	}
}

func TestAdapter_withEverything_Success(t *testing.T) {
	element := NewElementWithEverythingForTests()
	adapter := NewAdapter()
	content, err := adapter.ToContent(element)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retElement, err := adapter.ToElement(content)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !reflect.DeepEqual(element, retElement) {
		t.Errorf("the returned element is invalid")
		return
	}
}

func TestAdapter_withRecursive_Success(t *testing.T) {
	element := NewElementWithRecursiveForTests()
	adapter := NewAdapter()
	content, err := adapter.ToContent(element)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retElement, err := adapter.ToElement(content)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !reflect.DeepEqual(element, retElement) {
		t.Errorf("the returned element is invalid")
		return
	}
}
