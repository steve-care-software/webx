package developments

import "github.com/steve-care-software/webx/engine/bytes/domain/delimiters"

// Adapter represents a development adapter
type Adapter interface {
	InstancesToBytes(ins Developments) ([]byte, error)
	BytesToInstances(data []byte) (Developments, []byte, error)
	InstanceToBytes(ins Development) ([]byte, error)
	BytesToInstance(data []byte) (Development, []byte, error)
}

// Builder represents a developments builder
type Builder interface {
	Create() Builder
	WithList(list []Development) Builder
	Now() (Developments, error)
}

// Developments represents a developments instance
type Developments interface {
	List() []Development
}

// DevelopmentBuilder represents a development builder
type DevelopmentBuilder interface {
	Create() DevelopmentBuilder
	WithName(name string) DevelopmentBuilder
	WithDescription(description string) DevelopmentBuilder
	IsDeleted() DevelopmentBuilder
	WithBranches(branches delimiters.Delimiter) DevelopmentBuilder
	Now() (Development, error)
}

// Development represents a development iteration
type Development interface {
	Name() string
	Description() string
	IsDeleted() bool
	HasBranches() bool
	Branches() delimiters.Delimiter
}
