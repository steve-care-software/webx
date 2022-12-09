package fetchers

import "github.com/steve-care-software/webx/blockchains/domain/cryptography/hash"

type fetcher struct {
	hash    hash.Hash
	content Content
}

func createFetcher(
	hash hash.Hash,
	content Content,
) Fetcher {
	out := fetcher{
		hash:    hash,
		content: content,
	}

	return &out
}

// Hash returns the hash
func (obj *fetcher) Hash() hash.Hash {
	return obj.hash
}

// Content returns the content
func (obj *fetcher) Content() Content {
	return obj.content
}
