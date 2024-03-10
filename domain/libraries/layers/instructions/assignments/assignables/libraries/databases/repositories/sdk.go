package repositories

// Builder represents a repository builder
type Builder interface {
	Create() Builder
	WithList(list string) Builder
	WithRetrieve(retrieve string) Builder
	IsSkeleton() Builder
	IsHeight() Builder
	Now() (Repository, error)
}

// Repository represents a repository
type Repository interface {
	IsSkeleton() bool
	IsHeight() bool
	IsList() bool
	List() string
	IsRetrieve() bool
	Retrieve() string
}
