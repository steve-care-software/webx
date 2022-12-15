package blockchains

import "github.com/steve-care-software/webx/roots/domain/blockchains/cryptography/hash"

type content struct {
	hash hash.Hash
	data []byte
	kind uint
}
