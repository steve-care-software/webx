package pointers

import (
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/namespaces/singles/versions/singles/workspaces/singles/branches/singles/states/singles/pointers"
)

// Application represents a pointer application
type Application interface {
	Write(pointer pointers.Pointer) error
}
