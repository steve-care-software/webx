package iterations

import (
	"github.com/steve-care-software/webx/engine/bytes/domain/iterations/developments"
	"github.com/steve-care-software/webx/engine/bytes/domain/iterations/productions"
)

// Adapter represents an iteration adapter
type Adapter interface {
	InstancesToBytes(ins Iterations) ([]byte, error)
	BytesToInstances(data []byte) (Iterations, []byte, error)
	InstanceToBytes(ins Iteration) ([]byte, error)
	BytesToInstance(data []byte) (Iteration, []byte, error)
}

// Builder represents an iterations builder
type Builder interface {
	Create() Builder
	WithList(list []Iteration) Builder
	Now() (Iterations, error)
}

// Iterations represents iterations
type Iterations interface {
	List() []Iteration
}

// IterationBuilder represents an iteration builder
type IterationBuilder interface {
	Create() IterationBuilder
	WithName(name string) IterationBuilder
	WithDescription(description string) IterationBuilder
	IsDeleted() IterationBuilder
	WithDevelopments(developments developments.Developments) IterationBuilder
	WithProductions(productions productions.Productions) IterationBuilder
	Now() (Iteration, error)
}

// Iteration represents an iteration
type Iteration interface {
	Name() string
	Description() string
	IsDeleted() bool
	HasDevelopments() bool
	Developments() developments.Developments
	HasProductions() bool
	Productions() productions.Productions
}
