package updates

import "github.com/steve-care-software/webx/engine/domain/transpiles/blocks/lines/tokens/pointers"

// Update represents an update
type Update interface {
	Origin() pointers.Pointer
	Target() pointers.Pointer
}
