package queries

import (
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/databases/commits/queries/chunks"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// NewQueryBuilder creates a new query builder
func NewQueryBuilder() QueryBuilder {
	hashAdapter := hash.NewAdapter()
	return createQueryBuilder(
		hashAdapter,
	)
}

// Builder represents a queries builder
type Builder interface {
	Create() Builder
	WithList(list []Query) Builder
	Now() (Queries, error)
}

// Queries represents queries
type Queries interface {
	Hash() hash.Hash
	List() []Query
}

// QueryBuilder represents a query builder
type QueryBuilder interface {
	Create() QueryBuilder
	WithBytes(bytes []byte) QueryBuilder
	WithChunk(chunk chunks.Chunk) QueryBuilder
	Now() (Query, error)
}

// Query represents a query
type Query interface {
	Hash() hash.Hash
	IsBytes() bool
	Bytes() []byte
	IsChunk() bool
	Chunk() chunks.Chunk
}

// Repository represents a query repository
type Repository interface {
	Retrieve(hash hash.Hash) (Query, error)
	RetrieveAll(hashes []hash.Hash) (Queries, error)
}

// Service represents a query service
type Service interface {
	Save(ins Query) error
	SaveAll(list Queries) error
	Delete(hash hash.Hash) error
	DeleteAll(hashes []hash.Hash) error
}
