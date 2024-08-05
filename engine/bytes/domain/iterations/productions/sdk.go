package productions

import "github.com/steve-care-software/webx/engine/bytes/domain/delimiters"

// Adapter represents a production adapter
type Adapter interface {
	InstancesToBytes(ins Productions) ([]byte, error)
	BytesToInstances(data []byte) (Productions, []byte, error)
	InstanceToBytes(ins Production) ([]byte, error)
	BytesToInstance(data []byte) (Production, []byte, error)
}

// Builder represents a productions builder
type Builder interface {
	Create() Builder
	WithList(list []Production) Builder
	Now() (Productions, error)
}

// Productions represents a productions instance
type Productions interface {
	List() []Production
}

// ProductionBuilder represents a production builder
type ProductionBuilder interface {
	Create() ProductionBuilder
	WithName(name string) ProductionBuilder
	WithWorkspace(workspace delimiters.Delimiter) ProductionBuilder
	IsDeleted() ProductionBuilder
	Now() (Production, error)
}

// Production represents a production iteration
type Production interface {
	Name() string
	Description() string
	IsDeleted() bool
	Workspace() delimiters.Delimiter
}
