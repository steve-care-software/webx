package updates

import (
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/identities/switchers/singles/namespaces/switchers/singles/versions/switchers/singles"
)

// Builder represents an update builder
type Builder interface {
	Create() Builder
	WithSingle(single singles.Single) Builder
	WithBytes(bytes []byte) Builder
	Now() (Update, error)
}

// Update represents an update
type Update interface {
	Single() singles.Single
	Bytes() []byte
}
