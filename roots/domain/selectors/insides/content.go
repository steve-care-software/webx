package insides

import (
	"github.com/steve-care-software/webx/databases/domain/cryptography/hash"
)

type content struct {
	fn       hash.Hash
	fetchers []hash.Hash
}

func createContentWithFn(
	fn hash.Hash,
) Content {
	return createContentInternally(fn, nil)
}

func createContentWithFetchers(
	fetchers []hash.Hash,
) Content {
	return createContentInternally(nil, fetchers)
}

func createContentInternally(
	fn hash.Hash,
	fetchers []hash.Hash,
) Content {
	out := content{
		fn:       fn,
		fetchers: fetchers,
	}

	return &out
}

// IsFn returns true if there is a func, false otherwise
func (obj *content) IsFn() bool {
	return obj.fn != nil
}

// Fn returns the func, if any
func (obj *content) Fn() hash.Hash {
	return obj.fn
}

// IsFetchers returns true if there is a fetchers, false otherwise
func (obj *content) IsFetchers() bool {
	return obj.fetchers != nil
}

// Fetchers returns the fetchers, if any
func (obj *content) Fetchers() []hash.Hash {
	return obj.fetchers
}
