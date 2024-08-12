package transactions

import (
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/resources/storages/delimiters"
)

// Transaction represents a transaction
type Transaction interface {
	Index() uint64
	Transaction() delimiters.Delimiter
}
