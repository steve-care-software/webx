package elements

import (
	"reflect"
	"testing"
)

func TestCardinalityAdapter_Success(t *testing.T) {
	cardinality := NewCardinalityForTests(false)
	adapter := NewCardinalityAdapter()
	content, err := adapter.ToContent(cardinality)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retCardinality, err := adapter.ToCardinality(content)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !reflect.DeepEqual(cardinality, retCardinality) {
		t.Errorf("the returned cardinality is invalid")
		return
	}
}

func TestCardinalityAdapter_withMaximum_Success(t *testing.T) {
	cardinality := NewCardinalityForTests(true)
	adapter := NewCardinalityAdapter()
	content, err := adapter.ToContent(cardinality)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retCardinality, err := adapter.ToCardinality(content)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !reflect.DeepEqual(cardinality, retCardinality) {
		t.Errorf("the returned cardinality is invalid")
		return
	}
}
