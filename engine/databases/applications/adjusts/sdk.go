package applications

import (
	"github.com/steve-care-software/webx/engine/databases/domain/entries"
	"github.com/steve-care-software/webx/engine/databases/domain/headers"
	"github.com/steve-care-software/webx/engine/databases/domain/headers/containers"
)

// Application represents the adjust application
type Application interface {
	ExecuteWithInserts(original headers.Header, toInsert entries.Entries) (headers.Header, error)
	ExecuteWithDeletes(original headers.Header, toDelete containers.Containers) (headers.Header, error)
	ExecuteWithInsertsAndDeletes(original headers.Header, toInsert entries.Entries, toDelete containers.Containers) (headers.Header, error)
}
