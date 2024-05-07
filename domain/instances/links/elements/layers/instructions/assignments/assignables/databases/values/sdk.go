package values

import (
	"github.com/steve-care-software/datastencil/domain/hash"
)

// Value represents a value
type Value interface {
	Hash() hash.Hash
	IsInstance() bool
	Instance() string
	IsTransform() bool
	Transform() string
}
