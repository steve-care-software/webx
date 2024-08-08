package databases

import (
	"github.com/steve-care-software/webx/engine/cursors/domain/cursors"
	"github.com/steve-care-software/webx/engine/cursors/domain/hash"
)

// Database represents a database axtion
type Database interface {
	Hash() hash.Hash
	Cursor() cursors.Cursor
}
