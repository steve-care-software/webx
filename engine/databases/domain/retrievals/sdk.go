package retrievals

// Builder represents the retrievals builder
type Builder interface {
	Create() Builder
	WithList(list []Retrieval) Builder
	Now() (Retrievals, error)
}

// Retrievals represents retrievals
type Retrievals interface {
	List() []Retrieval
}

// RetrievalBuilder represents the retrieval builder
type RetrievalBuilder interface {
	Create() RetrievalBuilder
	WithIndex(index int64) RetrievalBuilder
	WithLength(length int64) RetrievalBuilder
	Now() (Retrieval, error)
}

// Retrieval represents a retrieval
type Retrieval interface {
	Index() int64
	Length() int64
}
