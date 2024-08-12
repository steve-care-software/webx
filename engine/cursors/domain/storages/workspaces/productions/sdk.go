package productions

import (
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/resources/storages/delimiters"
	"github.com/steve-care-software/webx/engine/cursors/domain/storages/originals"
)

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
	WithOriginal(original originals.Original) ProductionBuilder
	WithWorkspace(workspace delimiters.Delimiter) ProductionBuilder
	Now() (Production, error)
}

// Production represents a production iteration
type Production interface {
	Original() originals.Original
	Branch() delimiters.Delimiter
}
