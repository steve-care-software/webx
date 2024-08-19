package rules

import (
	"bytes"
	"testing"
)

func TestAdapter_withRule_Success(t *testing.T) {
	name := "MY_NAME"
	values := []byte("this is some bytes")
	rule := NewRuleForTests(name, values)
	adapter := NewAdapter()
	retNFT, err := adapter.RuleToNFT(rule)
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
