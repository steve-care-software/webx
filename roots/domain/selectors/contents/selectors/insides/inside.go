package insides

import "github.com/steve-care-software/webx/roots/domain/blockchains/cryptography/hash"

type inside struct {
	hash    hash.Hash
	content Content
}

func createInside(
	hash hash.Hash,
	content Content,
) Inside {
	out := inside{
		hash:    hash,
		content: content,
	}

	return &out
}

// Hash returns the hash
func (obj *inside) Hash() hash.Hash {
	return obj.hash
}

// Content returns the content
func (obj *inside) Content() Content {
	return obj.content
}
