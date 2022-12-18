package references

import (
	"reflect"
	"testing"
)

func TestCommitAdapter_Success(t *testing.T) {
	commit := NewCommitForTests()
	adapter := NewCommitAdapter()
	content, err := adapter.ToContent(commit)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retCommit, err := adapter.ToCommit(content)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !reflect.DeepEqual(commit, retCommit) {
		t.Errorf("the returned commit is invalid")
		return
	}
}
