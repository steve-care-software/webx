package grammars

import (
	"github.com/steve-care-software/logics/domain/bytes/grammars/cardinalities"
	"github.com/steve-care-software/logics/domain/bytes/grammars/values"
)

// Builder represents a grammar builder
type Builder interface {
	Create() Builder
	WithRoot(root Token) Builder
	WithChannels(channels Channels) Builder
	Now() (Grammar, error)
}

// Grammar represents a grammar
type Grammar interface {
	Root() Token
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
	List() []Channel
}

// ChannelBuilder represents a channel builder
type ChannelBuilder interface {
	Create() ChannelBuilder
	WithName(name string) ChannelBuilder
	WithToken(token Token) ChannelBuilder
	WithCondition(condition ChannelCondition) ChannelBuilder
	Now() (Channel, error)
}

// Channel represents a channel
type Channel interface {
	Name() string
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
	Name() string
	Grammar() Grammar
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
	IsValid() bool
	Valid() []byte
	IsInvalid() bool
	Invalid() []byte
}

// BlockBuilder represents a block builder
type BlockBuilder interface {
	Create() BlockBuilder
	WithLines(lines []Line) BlockBuilder
	Now() (Block, error)
}

// Block represents a decision block
type Block interface {
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
	Elements() []Element
}

// ElementBuilder represents an element builder
type ElementBuilder interface {
	Create() ElementBuilder
	WithName(name string) ElementBuilder
	WithCardinality(cardinality cardinalities.Cardinality) ElementBuilder
	WithValue(value values.Value) ElementBuilder
	WithToken(token Token) ElementBuilder
	WithExternal(external External) ElementBuilder
	Now() (Element, error)
}

// Element represents an element
type Element interface {
	Name() string
	Content() ElementContent
	Cardinality() cardinalities.Cardinality
}

// ElementContent represents an element content
type ElementContent interface {
	IsValue() bool
	Value() values.Value
	IsToken() bool
	Token() Token
	IsExternal() bool
	External() External
}
