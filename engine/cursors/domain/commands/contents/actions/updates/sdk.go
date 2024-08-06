package updates

import (
	"github.com/steve-care-software/webx/engine/cursors/domain/storages/originals"
)

// Update represents an update
type Update interface {
	Original() string
	Update() originals.Original
}
