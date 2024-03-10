package services

// Builder represents a service builder
type Builder interface {
	Create() Builder
	IsBegin() Builder
	Now() (Service, error)
}

// Service represents a service
type Service interface {
	IsBegin() bool
}
