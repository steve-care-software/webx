package files

import (
	"os"

	"github.com/juju/fslock"
	"github.com/steve-care-software/webx/engine/bytes/domain/delimiters"
	"github.com/steve-care-software/webx/engine/bytes/domain/entries"
	"github.com/steve-care-software/webx/engine/bytes/domain/states"
)

type context struct {
	path          []string
	name          string
	pDataIndex    *uint64
	currentHeader states.States
	insertions    entries.Entries
	deletions     delimiters.Delimiters
	pFile         *os.File
	pLock         *fslock.Lock
}
