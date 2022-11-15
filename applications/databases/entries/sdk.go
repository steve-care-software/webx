package entries

import (
	"github.com/steve-care-software/webx/domain/cryptography/hash"
	"github.com/steve-care-software/webx/domain/databases/entries"
)

// Application represents the pendings application
type Application interface {
	List(kind uint8) ([]byte, error)
	Retrieve(kind uint8, hash hash.Hash) ([]byte, error)
	Insert(entry entries.Entry) error
}
