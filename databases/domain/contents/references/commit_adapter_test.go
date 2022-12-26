package references

import (
	"reflect"
	"testing"
)

func TestCommitAdapter_Success(t *testing.T) {
	commit := NewCommitForTests()
	adapter := NewCommitAdapter([]byte("0")[0])
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

func TestCommitAdapter_withParent_Success(t *testing.T) {
	commit := NewCommitWithParentForTests()
	adapter := NewCommitAdapter([]byte("0")[0])
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

func TestCommitAdapter_withProof_Success(t *testing.T) {
	commit := NewCommitWithProofForTests()
	adapter := NewCommitAdapter([]byte("0")[0])
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

func TestCommitAdapter_withParent_withProof_Success(t *testing.T) {
	commit := NewCommitWithParentAndProofForTests()
	adapter := NewCommitAdapter([]byte("0")[0])
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
