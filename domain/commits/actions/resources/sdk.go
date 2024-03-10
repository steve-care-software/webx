package resources

import (
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/commits/actions/resources/instances"
)

// Resources represents resources
type Resources interface {
	Hash() hash.Hash
	List() []Resource
}

// Resource represents a resource
type Resource interface {
	Hash() hash.Hash
	Path() []string
	Instance() instances.Instance
}
