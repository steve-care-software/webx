package databases

import "github.com/steve-care-software/datastencil/domain/skeletons"

// Builder represents a database builder
type Builder interface {
	Create() Builder
	WithSkeleton(skeleton skeletons.Skeleton) Builder
	Now() (Database, error)
}

// Database represents a database
type Database interface {
	IsSkeleton() bool
	Skeleton() skeletons.Skeleton
}
