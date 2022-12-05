package insides

import "github.com/steve-care-software/webx/domain/databases/entities"

type content struct {
	fn       entities.Identifier
	fetchers entities.Identifiers
}

func createContentWithFn(
	fn entities.Identifier,
) Content {
	return createContentInternally(fn, nil)
}

func createContentWithFetchers(
	fetchers entities.Identifiers,
) Content {
	return createContentInternally(nil, fetchers)
}

func createContentInternally(
	fn entities.Identifier,
	fetchers entities.Identifiers,
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
func (obj *content) Fn() entities.Identifier {
	return obj.fn
}

// IsFetchers returns true if there is a fetchers, false otherwise
func (obj *content) IsFetchers() bool {
	return obj.fetchers != nil
}

// Fetchers returns the fetchers, if any
func (obj *content) Fetchers() entities.Identifiers {
	return obj.fetchers
}
