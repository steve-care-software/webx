package executions

// Builder represents an execution builder
type Builder interface {
	Create() Builder
	WithCredentials(credentials string) Builder
	WithAccount(account string) Builder
	WithCriteria(criteria string) Builder
	Now() (Execution, error)
}

// Execution represents an update execution
type Execution interface {
	Credentials() string
	Account() string
	Criteria() string
}
