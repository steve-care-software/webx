package listers

import "github.com/steve-care-software/webx/engine/databases/bytes/domain/retrievals"

type lister struct {
	keyname   string
	retrieval retrievals.Retrieval
}

func createLister(
	keyname string,
	retrieval retrievals.Retrieval,
) Lister {
	out := lister{
		keyname:   keyname,
		retrieval: retrieval,
	}

	return &out
}

// Keyname returns the keyname
func (obj *lister) Keyname() string {
	return obj.keyname
}

// Retrieval returns the retrieval
func (obj *lister) Retrieval() retrievals.Retrieval {
	return obj.retrieval
}
