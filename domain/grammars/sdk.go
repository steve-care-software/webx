package grammars

import (
	"github.com/steve-care-software/webx/domain/cryptography/hash"
	"github.com/steve-care-software/webx/domain/grammars/cardinalities"
	"github.com/steve-care-software/webx/domain/grammars/values"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(hashAdapter)
}

// NewChannelsBuilder creates a new channels builder
func NewChannelsBuilder() ChannelsBuilder {
	hashAdapter := hash.NewAdapter()
	return createChannelsBuilder(hashAdapter)
}

// NewChannelBuilder creates a new channel builder
func NewChannelBuilder() ChannelBuilder {
	hashAdapter := hash.NewAdapter()
	return createChannelBuilder(hashAdapter)
}

// NewChannelConditionBuilder creates a new chanel condition builder
func NewChannelConditionBuilder() ChannelConditionBuilder {
	hashAdapter := hash.NewAdapter()
	return createChannelConditionBuilder(hashAdapter)
}

// NewExternalBuilder creates a new external builder
func NewExternalBuilder() ExternalBuilder {
	hashAdapter := hash.NewAdapter()
	return createExternalBuilder(hashAdapter)
}

// NewInstanceBuilder creates a new instance builder
func NewInstanceBuilder() InstanceBuilder {
	return createInstanceBuilder()
}

// NewEverythingBuilder creates a new everything builder
func NewEverythingBuilder() EverythingBuilder {
	hashAdapter := hash.NewAdapter()
	return createEverythingBuilder(hashAdapter)
}

// NewTokensBuilder creates a new tokens builder
func NewTokensBuilder() TokensBuilder {
	hashAdapter := hash.NewAdapter()
	return createTokensBuilder(hashAdapter)
}

// NewTokenBuilder creates a new token builder
func NewTokenBuilder() TokenBuilder {
	hashAdapter := hash.NewAdapter()
	return createTokenBuilder(hashAdapter)
}

// NewSuitesBuilder creates a new suites builder
func NewSuitesBuilder() SuitesBuilder {
	hashAdapter := hash.NewAdapter()
	return createSuitesBuilder(hashAdapter)
}

// NewSuiteBuilder creates a new suite builder
func NewSuiteBuilder() SuiteBuilder {
	hashAdapter := hash.NewAdapter()
	return createSuiteBuilder(hashAdapter)
}

// NewBlockBuilder creates a new block builder
func NewBlockBuilder() BlockBuilder {
	hashAdapter := hash.NewAdapter()
	return createBlockBuilder(hashAdapter)
}

// NewLineBuilder creates a new line builder
func NewLineBuilder() LineBuilder {
	hashAdapter := hash.NewAdapter()
	return createLineBuilder(hashAdapter)
}

// NewElementBuilder creates a new element builder
func NewElementBuilder() ElementBuilder {
	hashAdapter := hash.NewAdapter()
	return createElementBuilder(hashAdapter)
}

// Builder represents a grammar builder
type Builder interface {
	Create() Builder
	WithRoot(root Token) Builder
	WithChannels(channels Channels) Builder
	Now() (Grammar, error)
}

// Grammar represents a grammar
type Grammar interface {
	Hash() hash.Hash
	Root() Token
	HasChannels() bool
	Channels() Channels
}

// ChannelsBuilder represents a channels builder
type ChannelsBuilder interface {
	Create() ChannelsBuilder
	WithList(list []Channel) ChannelsBuilder
	Now() (Channels, error)
}

// Channels represents channels
type Channels interface {
	Hash() hash.Hash
	List() []Channel
}

// ChannelBuilder represents a channel builder
type ChannelBuilder interface {
	Create() ChannelBuilder
	WithToken(token Token) ChannelBuilder
	WithCondition(condition ChannelCondition) ChannelBuilder
	Now() (Channel, error)
}

// Channel represents a channel
type Channel interface {
	Hash() hash.Hash
	Token() Token
	HasCondition() bool
	Condition() ChannelCondition
}

// ChannelConditionBuilder represents a channel condition builder
type ChannelConditionBuilder interface {
	Create() ChannelConditionBuilder
	WithPrevious(previous Token) ChannelConditionBuilder
	WithNext(next Token) ChannelConditionBuilder
	Now() (ChannelCondition, error)
}

// ChannelCondition represents a channel condition
type ChannelCondition interface {
	Hash() hash.Hash
	HasPrevious() bool
	Previous() Token
	HasNext() bool
	Next() Token
}

// ExternalBuilder represents an external builder
type ExternalBuilder interface {
	Create() ExternalBuilder
	WithName(name string) ExternalBuilder
	WithGrammar(grammar Grammar) ExternalBuilder
	Now() (External, error)
}

// External represents an external token
type External interface {
	Hash() hash.Hash
	Name() string
	Grammar() Grammar
}

// InstanceBuilder represents an instance builder
type InstanceBuilder interface {
	Create() InstanceBuilder
	WithToken(token Token) InstanceBuilder
	WithEverything(everything Everything) InstanceBuilder
	Now() (Instance, error)
}

// Instance represents an instance
type Instance interface {
	Hash() hash.Hash
	Name() string
	IsToken() bool
	Token() Token
	IsEverything() bool
	Everything() Everything
}

// EverythingBuilder represents an everything builder
type EverythingBuilder interface {
	Create() EverythingBuilder
	WithName(name string) EverythingBuilder
	WithException(exception Token) EverythingBuilder
	WithEscape(escape Token) EverythingBuilder
	Now() (Everything, error)
}

// Everything represents an everything except
type Everything interface {
	Hash() hash.Hash
	Name() string
	Exception() Token
	HasEscape() bool
	Escape() Token
}

// TokensBuilder represents a tokens builder
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
	WithName(name string) TokenBuilder
	WithBlock(block Block) TokenBuilder
	WithSuites(suites Suites) TokenBuilder
	Now() (Token, error)
}

// Token represents a token
type Token interface {
	Hash() hash.Hash
	Name() string
	Block() Block
	HasSuites() bool
	Suites() Suites
}

// SuitesBuilder represents a suites builder
type SuitesBuilder interface {
	Create() SuitesBuilder
	WithList(list []Suite) SuitesBuilder
	Now() (Suites, error)
}

// Suites represets a list of test suites
type Suites interface {
	Hash() hash.Hash
	List() []Suite
}

// SuiteBuilder represents a suite builder
type SuiteBuilder interface {
	Create() SuiteBuilder
	WithValid(valid []byte) SuiteBuilder
	WithInvalid(invalid []byte) SuiteBuilder
	Now() (Suite, error)
}

// Suite represents a test suite
type Suite interface {
	Hash() hash.Hash
	IsValid() bool
	Content() []byte
}

// BlockBuilder represents a block builder
type BlockBuilder interface {
	Create() BlockBuilder
	WithLines(lines []Line) BlockBuilder
	Now() (Block, error)
}

// Block represents a decision block
type Block interface {
	Hash() hash.Hash
	Lines() []Line
}

// LineBuilder represents a line builder
type LineBuilder interface {
	Create() LineBuilder
	WithElements(elements []Element) LineBuilder
	Now() (Line, error)
}

// Line represents a line of elements
type Line interface {
	Hash() hash.Hash
	Elements() []Element
}

// ElementBuilder represents an element builder
type ElementBuilder interface {
	Create() ElementBuilder
	WithCardinality(cardinality cardinalities.Cardinality) ElementBuilder
	WithValue(value values.Value) ElementBuilder
	WithExternal(external External) ElementBuilder
	WithInstance(instance Instance) ElementBuilder
	WithRecursive(recursive string) ElementBuilder
	Now() (Element, error)
}

// Element represents an element
type Element interface {
	Hash() hash.Hash
	Name() string
	Content() ElementContent
	Cardinality() cardinalities.Cardinality
}

// ElementContent represents an element content
type ElementContent interface {
	Hash() hash.Hash
	IsValue() bool
	Value() values.Value
	IsExternal() bool
	External() External
	IsInstance() bool
	Instance() Instance
	IsRecursive() bool
	Recursive() string
}
