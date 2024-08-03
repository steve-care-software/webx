package routes

import (
	"github.com/steve-care-software/webx/engine/hashes/domain/hash"
	"github.com/steve-care-software/webx/engine/vms/domain/instances/routes/omissions"
	"github.com/steve-care-software/webx/engine/vms/domain/instances/routes/tokens"
)

// NewBuilder creates a new route builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// Adapter represents the route adapter
type Adapter interface {
	ToBytes(ins Route) ([]byte, error)
	ToInstance(bytes []byte) (Route, error)
}

// Builder represents a route builder
type Builder interface {
	Create() Builder
	WithLayer(layer hash.Hash) Builder
	WithTokens(tokens tokens.Tokens) Builder
	WithGlobal(omission omissions.Omission) Builder
	WithToken(token omissions.Omission) Builder
	Now() (Route, error)
}

// Route represents a route
type Route interface {
	Hash() hash.Hash
	Layer() hash.Hash
	Tokens() tokens.Tokens
	HasGlobal() bool
	Global() omissions.Omission
	HasToken() bool
	Token() omissions.Omission
}

// Repository represents a route repository
type Repository interface {
	Amount() (*uint, error)
	List(index uint, amount uint) ([]hash.Hash, error)
	Retrieve(hash hash.Hash) (Route, error)
	RetrieveFromLayer(layerHash hash.Hash) (Route, error)
}
