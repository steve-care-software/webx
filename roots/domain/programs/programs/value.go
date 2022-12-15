package programs

import "github.com/steve-care-software/webx/roots/domain/blockchains/cryptography/hash"

type value struct {
	hash    hash.Hash
	content Content
}

func createValue(
	hash hash.Hash,
	content Content,
) Value {
	out := value{
		hash:    hash,
		content: content,
	}

	return &out
}

// Hash returns the hash
func (obj *value) Hash() hash.Hash {
	return obj.hash
}

// Content returns the content
func (obj *value) Content() Content {
	return obj.content
}
