package applications

import (
	"github.com/steve-care-software/webx/engine/hashes/domain/hash"
	"github.com/steve-care-software/webx/engine/hashes/domain/pointers"
)

type context struct {
	current pointers.Pointers
	inserts []pointers.Pointer
	deletes []hash.Hash
}
