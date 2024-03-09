package transactions

// Transaction represents a transaction
type Transaction interface {
	IsResource() bool
	Resource() Resource
	IsPointer() bool
	Pointer() Pointer
}

// Resource represents a resource
type Resource interface {
	Path() string
	Instance() string
}

// Pointer represents a pointer
type Pointer interface {
	Path() string
	Hash() string
}
