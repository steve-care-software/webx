package delimiters

// NewBuilder creates a new builder
func NewBuilder() Builder {
	return createBuilder()
}

// NewDelimiterBuilder creates a new delimiter builder
func NewDelimiterBuilder() DelimiterBuilder {
	return createDelimiterBuilder()
}

// Adapter represents a delimiter adapter
type Adapter interface {
	InstancesToBytes(ins Delimiters) ([]byte, error)
	BytesToInstances(data []byte) (Delimiters, error)
	InstanceToBytes(ins Delimiter) ([]byte, error)
	BytesToInstance(data []byte) (Delimiter, error)
}

// Builder represents the delimiters builder
type Builder interface {
	Create() Builder
	WithList(list []Delimiter) Builder
	Now() (Delimiters, error)
}

// Delimiters represents delimiters
type Delimiters interface {
	List() []Delimiter
}

// DelimiterBuilder represents the delimiter builder
type DelimiterBuilder interface {
	Create() DelimiterBuilder
	WithIndex(index uint64) DelimiterBuilder
	WithLength(length uint64) DelimiterBuilder
	Now() (Delimiter, error)
}

// Delimiter represents a delimiter
type Delimiter interface {
	Index() uint64
	Length() uint64
}
