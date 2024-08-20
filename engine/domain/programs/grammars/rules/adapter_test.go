package rules

import (
	"bytes"
	"testing"
)

func TestAdapter_withRules_Success(t *testing.T) {
	rules := NewRulesForTests([]Rule{
		NewRuleForTests("FIRST", []byte("first values")),
		NewRuleForTests("SECOND", []byte("second values")),
	})

	adapter := NewAdapter()
	retNFT, err := adapter.InstancesToNFT(rules)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retRules, err := adapter.NFTToInstances(retNFT)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	length := len(retRules.List())
	if length != 2 {
		t.Errorf("the list was expected to contain %d Rule instances, %d returned", 2, length)
		return
	}
}

func TestAdapter_withRule_Success(t *testing.T) {
	name := "MY_NAME"
	values := []byte("this is some bytes")
	rule := NewRuleForTests(name, values)
	adapter := NewAdapter()
	retNFT, err := adapter.InstanceToNFT(rule)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retRule, err := adapter.NFTToInstance(retNFT)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retName := retRule.Name()
	if retName != name {
		t.Errorf("the name was expected to be (%s), %s returned", name, retName)
		return
	}

	retValues := retRule.Bytes()
	if !bytes.Equal(values, retValues) {
		t.Errorf("the returned bytes are invalid")
		return
	}
}
