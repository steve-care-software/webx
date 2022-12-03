package cardinalities

import (
	"reflect"
	"testing"
)

func TestAdapter_Success(t *testing.T) {
	cardinality := NewCardinalityForTests(false)
	adapter := NewAdapter()
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

func TestAdapter_withMaximum_Success(t *testing.T) {
	cardinality := NewCardinalityForTests(true)
	adapter := NewAdapter()
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
