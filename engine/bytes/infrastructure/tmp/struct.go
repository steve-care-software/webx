package tmp

import (
	"os"

	"github.com/juju/fslock"
	"github.com/steve-care-software/webx/engine/bytes/domain/namespaces"
)

type context struct {
	dbPath              []string
	dbName              string
	namespaces          namespaces.Namespaces
	pNamespaceDataIndex *uint64
	currentNamespace    namespaces.Namespace
	pFile               *os.File
	pLock               *fslock.Lock
}
