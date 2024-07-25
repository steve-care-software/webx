package routes

import (
	"github.com/steve-care-software/webx/engine/states/domain/hash"
	"github.com/steve-care-software/webx/engine/vms/domain/instances/routes/cardinalities"
)

// NewElementsBuilder creates a new elements builder
func NewElementsBuilder() ElementsBuilder {
	hashAdapter := hash.NewAdapter()
	return createElementsBuilder(
		hashAdapter,
	)
}

// NewElementBuilder creates a new element builder
func NewElementBuilder() ElementBuilder {
	hashAdapter := hash.NewAdapter()
	return createElementBuilder(
		hashAdapter,
	)
}

// Builder represents a route builder
type Builder interface {
	Create() Builder
	WithTokens(tokens Tokens) Builder
	WithGlobal(omission Omission) Builder
	WithToken(token Omission) Builder
	Now() (Route, error)
}

// Route represents a route
type Route interface {
	Hash() hash.Hash
	Tokens() Tokens
	HasGlobal() bool
	Global() Omission
	HasToken() bool
	Token() Omission
}

// Repository represents a route repository
type Repository interface {
	List(index uint, amount uint) ([]hash.Hash, error)
	Retrieve(hash hash.Hash) (Route, error)
}

// Service represents a route service
type Service interface {
	Insert(ins Route) error
	Delete(hash hash.Hash) error
}

// TokensBuilder represents tokens builder
type TokensBuilder interface {
	Create() TokensBuilder
	WithList(list []Token) TokensBuilder
	Now() (Tokens, error)
}

// Tokens represents tokens
type Tokens interface {
	Hash() hash.Hash
	List() []Token
}

// TokenBuilder represents a token builder
type TokenBuilder interface {
	Create() TokenBuilder
	WithElements(elements Elements) TokenBuilder
	WithCardinality(cardinality cardinalities.Cardinality) TokenBuilder
	WithOmission(omission Omission) TokenBuilder
	Now() (Token, error)
}

// Token represents a token
type Token interface {
	Hash() hash.Hash
	Elements() Elements
	Cardinality() cardinalities.Cardinality
	HasOmission() bool
	Omission() Omission
}

// OmissionBuilder represents the omission builder
type OmissionBuilder interface {
	Create() OmissionBuilder
	WithPrefix(prefix Element) OmissionBuilder
	WithSuffix(suffix Element) OmissionBuilder
	Now() (Omission, error)
}

// Omission represents an omission
type Omission interface {
	Hash() hash.Hash
	HasPrefix() bool
	Prefix() Element
	HasSuffix() bool
	Suffix() Element
}

// ElementsBuilder represents the elements builder
type ElementsBuilder interface {
	Create() ElementsBuilder
	WithList(list []Element) ElementsBuilder
	Now() (Elements, error)
}

// Elements represents elements
type Elements interface {
	Hash() hash.Hash
	List() []Element
}

// ElementBuilder represents an element builder
type ElementBuilder interface {
	Create() ElementBuilder
	WithLayer(layer []string) ElementBuilder
	WithBytes(bytes []byte) ElementBuilder
	WithString(str string) ElementBuilder
	Now() (Element, error)
}

// Element represents a route element
type Element interface {
	Hash() hash.Hash
	IsLayer() bool
	Layer() []string
	IsBytes() bool
	Bytes() []byte
	IsString() bool
	String() string
}
