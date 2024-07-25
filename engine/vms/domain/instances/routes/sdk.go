package routes

import (
	"github.com/steve-care-software/webx/engine/states/domain/hash"
	"github.com/steve-care-software/webx/engine/vms/domain/instances/routes/cardinalities"
)

// NewBuilder creates a new route builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// NewTokensBuilder creates a new tokens builder
func NewTokensBuilder() TokensBuilder {
	hashAdapter := hash.NewAdapter()
	return createTokensBuilder(
		hashAdapter,
	)
}

// NewTokenBuilder creates a new token builder
func NewTokenBuilder() TokenBuilder {
	hashAdapter := hash.NewAdapter()
	return createTokenBuilder(
		hashAdapter,
	)
}

// NewOmissionBuilder creates a new omission builder
func NewOmissionBuilder() OmissionBuilder {
	hashAdapter := hash.NewAdapter()
	return createOmissionBuilder(
		hashAdapter,
	)
}

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
	WithLayer(layer hash.Hash) Builder
	WithTokens(tokens Tokens) Builder
	WithGlobal(omission Omission) Builder
	WithToken(token Omission) Builder
	Now() (Route, error)
}

// Route represents a route
type Route interface {
	Hash() hash.Hash
	Remaining(input []byte) []byte
	Layer() hash.Hash
	Tokens() Tokens
	HasGlobal() bool
	Global() Omission
	HasToken() bool
	Token() Omission
}

// Repository represents a route repository
type Repository interface {
	Amount() (*uint, error)
	List(index uint, amount uint) ([]hash.Hash, error)
	Retrieve(hash hash.Hash) (Route, error)
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
	RemainingWithOmission(input []byte, tokenOmission Omission) []byte
	Remaining(input []byte) []byte
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
	Remaining(input []byte) []byte
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
	Remaining(input []byte) []byte
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
	WithLayer(layer hash.Hash) ElementBuilder
	WithBytes(bytes []byte) ElementBuilder
	WithString(str string) ElementBuilder
	Now() (Element, error)
}

// Element represents a route element
type Element interface {
	Hash() hash.Hash
	IsLayer() bool
	Layer() hash.Hash
	IsBytes() bool
	Bytes() []byte
	IsString() bool
	String() string
}
