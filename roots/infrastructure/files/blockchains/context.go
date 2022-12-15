package blockchains

import (
	"os"

	"github.com/steve-care-software/webx/roots/domain/blockchains/contents/references"
)

type context struct {
	identifier  uint
	pConn       *os.File
	reference   references.Reference
	contentList []*content
}
