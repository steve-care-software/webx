package files

import "github.com/steve-care-software/webx/databases/domain/cryptography/hash"

type content struct {
	hash hash.Hash
	data []byte
	kind uint
}
