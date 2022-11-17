package modifications

import "github.com/steve-care-software/webx/domain/cryptography/hash"

type modification struct {
	hash    hash.Hash
	content Content
}

func createModification(
	hash hash.Hash,
	content Content,
) Modification {
	out := modification{
		hash:    hash,
		content: content,
	}

	return &out
}

// Hash returns the hash
func (obj *modification) Hash() hash.Hash {
	return obj.hash
}

// Content returns the content
func (obj *modification) Content() Content {
	return obj.content
}
