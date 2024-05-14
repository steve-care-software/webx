package commits

import "github.com/steve-care-software/datastencil/domain/hash"

type commit struct {
	hash    hash.Hash
	content Content
	parent  hash.Hash
}

func createCommit(
	hash hash.Hash,
	content Content,
) Commit {
	return createCommitInternally(hash, content, nil)
}

func createCommitWithParent(
	hash hash.Hash,
	content Content,
	parent hash.Hash,
) Commit {
	return createCommitInternally(hash, content, parent)
}

func createCommitInternally(
	hash hash.Hash,
	content Content,
	parent hash.Hash,
) Commit {
	out := commit{
		hash:    hash,
		content: content,
		parent:  parent,
	}

	return &out
}

// Hash returns the hash
func (obj *commit) Hash() hash.Hash {
	return obj.hash
}

// Content returns the content
func (obj *commit) Content() Content {
	return obj.content
}

// HasParent returns true if there is a parent, false otherwise
func (obj *commit) HasParent() bool {
	return obj.parent != nil
}

// Parent returns the parent, if any
func (obj *commit) Parent() hash.Hash {
	return obj.parent
}
