package references

import (
	"reflect"
	"testing"
)

func TestCommitsAdapter_Success(t *testing.T) {
	commits := NewCommitsForTests(25)
	adapter := NewCommitsAdapter([]byte("0")[0])
	content, err := adapter.ToContent(commits)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retCommits, err := adapter.ToCommits(content)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !reflect.DeepEqual(commits, retCommits) {
		t.Errorf("the returned commits is invalid")
		return
	}
}
