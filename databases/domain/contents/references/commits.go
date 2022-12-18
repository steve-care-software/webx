package references

import (
	"errors"
	"fmt"

	"github.com/steve-care-software/webx/databases/domain/cryptography/hash"
)

type commits struct {
	mp   map[string]Commit
	list []Commit
}

func createCommits(
	mp map[string]Commit,
	list []Commit,
) Commits {
	out := commits{
		mp:   mp,
		list: list,
	}

	return &out
}

// List returns the commits
func (obj *commits) List() []Commit {
	return obj.list
}

// Latest returns the latest commit
func (obj *commits) Latest() Commit {
	return obj.list[len(obj.list)-1]
}

// Fetch fetches a commit by hash
func (obj *commits) Fetch(hash hash.Hash) (Commit, error) {
	commitname := hash.String()
	if ins, ok := obj.mp[commitname]; ok {
		return ins, nil
	}

	str := fmt.Sprintf("the commit (hash: %s) is invalid", commitname)
	return nil, errors.New(str)
}
