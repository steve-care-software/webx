package transforms

import (
	"github.com/steve-care-software/datastencil/domain/hash"
)

// Transform represents a transform
type Transform interface {
	Hash() hash.Hash
	Query() string
	Bytes() string
}
