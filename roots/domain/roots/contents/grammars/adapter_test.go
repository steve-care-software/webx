package grammars

import (
	"reflect"
	"testing"
)

func TestAdapter_Success(t *testing.T) {
	grammar := NewGrammarForTests()
	adapter := NewAdapter()
	content, err := adapter.ToContent(grammar)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retGrammar, err := adapter.ToGrammar(content)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !reflect.DeepEqual(grammar, retGrammar) {
		t.Errorf("the returned grammar is invalid")
		return
	}
}

func TestAdapter_withHistory_Success(t *testing.T) {
	grammar := NewGrammarWithHistoryForTests()
	adapter := NewAdapter()
	content, err := adapter.ToContent(grammar)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retGrammar, err := adapter.ToGrammar(content)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !reflect.DeepEqual(grammar, retGrammar) {
		t.Errorf("the returned grammar is invalid")
		return
	}
}
