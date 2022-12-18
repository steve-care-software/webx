package files

import (
	"os"

	"github.com/steve-care-software/webx/databases/domain/contents/references"
)

type context struct {
	identifier  uint
	pConn       *os.File
	reference   references.Reference
	contentList []*content
}
