package headers

import "github.com/steve-care-software/webx/engine/cursors/domain/loaders/namespaces/singles/versions/singles/workspaces/singles/branches/singles/states/singles/pointers/storages"

// NewHeaderWithIdentitiesForTests creates a new header with identities for tests
func NewHeaderWithIdentitiesForTests(identities storages.Storage) Header {
	ins, err := NewBuilder().Create().WithIdentities(identities).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
