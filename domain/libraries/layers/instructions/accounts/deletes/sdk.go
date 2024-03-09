package deletes

// Builder represents a delete builder
type Builder interface {
	Create() Builder
	WithCredentials(credentials string) Builder
	WithAccount(account string) Builder
	Now() (Delete, error)
}

// Delete represents a delete
type Delete interface {
	Credentials() string
	Account() string
}
