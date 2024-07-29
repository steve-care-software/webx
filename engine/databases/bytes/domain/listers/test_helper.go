package listers

import "github.com/steve-care-software/webx/engine/databases/bytes/domain/retrievals"

// NewListerForTests creates a new lister for tests
func NewListerForTests(keyname string, retrieval retrievals.Retrieval) Lister {
	ins, err := NewBuilder().Create().WithKeyname(keyname).WithRetrieval(retrieval).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
