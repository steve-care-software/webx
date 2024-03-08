package orms

import (
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/orms"
	"github.com/steve-care-software/datastencil/domain/orms/skeletons"
)

// Application represents the ORM application
type Application interface {
	Skeleton() skeletons.Skeleton
	Retrieve(path []string, hash hash.Hash) (orms.Instance, error)
	Insert(resources orms.Resources) error
	Delete(resources orms.Resources) error
}
