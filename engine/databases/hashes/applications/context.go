package applications

import (
	"github.com/steve-care-software/webx/engine/databases/entities/domain/hash"
	"github.com/steve-care-software/webx/engine/databases/hashes/domain/pointers"
)

type context struct {
	current pointers.Pointers
	inserts []pointers.Pointer
	deletes []hash.Hash
}
