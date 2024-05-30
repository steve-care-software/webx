package conditions

import (
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/pointers/conditions/operators"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// NewResourceBuilder creates a new resource builder
func NewResourceBuilder() ResourceBuilder {
	hashAdapter := hash.NewAdapter()
	return createResourceBuilder(
		hashAdapter,
	)
}

// NewComparisonsBuilder creates a new comparisons builder
func NewComparisonsBuilder() ComparisonsBuilder {
	hashAdapter := hash.NewAdapter()
	return createComparisonsBuilder(
		hashAdapter,
	)
}

// NewComparisonBuilder creates a new comparison builder
func NewComparisonBuilder() ComparisonBuilder {
	hashAdapter := hash.NewAdapter()
	return createComparisonBuilder(
		hashAdapter,
	)
}

// Builder represents a condition builder
type Builder interface {
	Create() Builder
	WithResource(resource Resource) Builder
	WithComparisons(comparisons Comparisons) Builder
	Now() (Condition, error)
}

// Condition represents a condition
type Condition interface {
	Hash() hash.Hash
	Resource() Resource
	HasComparisons() bool
	Comparisons() Comparisons
}

// ResourceBuilder represents a resource builder
type ResourceBuilder interface {
	Create() ResourceBuilder
	WithPath(path []string) ResourceBuilder
	MustBeLoaded() ResourceBuilder
	Now() (Resource, error)
}

// Resource represents a resource
type Resource interface {
	Hash() hash.Hash
	Path() []string
	MustBeLoaded() bool
}

// ComparisonsBuilder represents a comparisons builder
type ComparisonsBuilder interface {
	Create() ComparisonsBuilder
	WithList(list []Comparison) ComparisonsBuilder
	Now() (Comparisons, error)
}

// Comparisons represents comparisons
type Comparisons interface {
	Hash() hash.Hash
	List() []Comparison
}

// ComparisonBuilder represents a comparison builder
type ComparisonBuilder interface {
	Create() ComparisonBuilder
	WithOperator(operator operators.Operator) ComparisonBuilder
	WithCondition(condition Condition) ComparisonBuilder
	Now() (Comparison, error)
}

// Comparison represents a comparison
type Comparison interface {
	Hash() hash.Hash
	Operator() operators.Operator
	Condition() Condition
}
