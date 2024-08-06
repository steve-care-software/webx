package updates

import "github.com/steve-care-software/webx/engine/cursors/domain/storages/delimiters"

// Update represents an update
type Update interface {
	Original() delimiters.Delimiter
	Update() []byte
}
