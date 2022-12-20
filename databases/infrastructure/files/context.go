package files

import (
	"net/url"
	"os"

	"github.com/juju/fslock"
	"github.com/steve-care-software/webx/databases/domain/connections/contents"
	"github.com/steve-care-software/webx/databases/domain/contents/references"
)

type context struct {
	identifier  uint
	name        string
	pLock       *fslock.Lock
	pConn       *os.File
	reference   references.Reference
	dataOffset  uint
	contentList []contents.Content
	peerList    []*url.URL
}
