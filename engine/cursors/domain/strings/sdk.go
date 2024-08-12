package strings

// Adapter represents a list adapter
type Adapter interface {
	ToBytes(ins Strings) ([]byte, error)
	ToInstance(data []byte) (Strings, error)
}

// Builder represents a list of string
type Builder interface {
	Create() Builder
	WithList(list []string) Builder
	IsUnique() Builder
	Now() (Strings, error)
}

// Strings represents a list of string
type Strings interface {
	IsUnique() bool
	List() []string
}
