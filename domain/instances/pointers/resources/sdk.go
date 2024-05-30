package resources

import (
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/databases/heads"
	"github.com/steve-care-software/datastencil/domain/instances/pointers/resources/logics"
)

// Builder represents a resources builder
type Builder interface {
	Create() Builder
	WithList(list []Resource) Builder
	Now() (Resources, error)
}

// Resources represents resources
type Resources interface {
	Hash() hash.Hash
	List() []Resource
}

// ResourceBuilder represents a resource builder
type ResourceBuilder interface {
	Create() ResourceBuilder
	WithDatabase(database heads.Head) ResourceBuilder
	WithLogics(logics logics.Logics) ResourceBuilder
	Now() (Resource, error)
}

// Resource represents a resource
type Resource interface {
	Hash() hash.Hash
	Database() heads.Head
	Logics() logics.Logics
}
