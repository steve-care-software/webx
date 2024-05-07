package actions

import "github.com/steve-care-software/datastencil/domain/hash"

type action struct {
	hash    hash.Hash
	path    []string
	content Content
}

func createAction(
	hash hash.Hash,
	path []string,
	content Content,
) Action {
	out := action{
		hash:    hash,
		path:    path,
		content: content,
	}

	return &out
}

// Hash returns the hash
func (obj *action) Hash() hash.Hash {
	return obj.hash
}

// Path returns the path
func (obj *action) Path() []string {
	return obj.path
}

// Content returns the content
func (obj *action) Content() Content {
	return obj.content
}
