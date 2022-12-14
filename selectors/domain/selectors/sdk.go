package selectors

import (
	"github.com/steve-care-software/webx/grammars/domain/grammars"
	"github.com/steve-care-software/webx/grammars/domain/trees"
)

// MultiContentFn represents the multi content func
type MultiContentFn func(contents []trees.Content) ([]interface{}, error)

// SingleContentFn represents the single content func
type SingleContentFn func(content trees.Content) ([]interface{}, error)

// MultiSelectorFn represents the multi selector func
type MultiSelectorFn func(instances []interface{}) (interface{}, bool, error)

// SingleSelectorFn represents the single selector func
type SingleSelectorFn func(instance interface{}) (interface{}, bool, error)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// NewSelectorFnBuilder creates a new selectorFn builder
func NewSelectorFnBuilder() SelectorFnBuilder {
	return createSelectorFnBuilder()
}

// NewTokenBuilder creates a new token builder
func NewTokenBuilder() TokenBuilder {
	return createTokenBuilder()
}

// NewElementBuilder creates a new element builder
func NewElementBuilder() ElementBuilder {
	return createElementBuilder()
}

// NewInsideBuilder creates a new inside builder
func NewInsideBuilder() InsideBuilder {
	return createInsideBuilder()
}

// NewFetchersBuilder creates a new fetchers builder
func NewFetchersBuilder() FetchersBuilder {
	return createFetchersBuilder()
}

// NewFetcherBuilder creates a new fetcher builder
func NewFetcherBuilder() FetcherBuilder {
	return createFetcherBuilder()
}

// NewContentFnBuilder creates a new content func builder
func NewContentFnBuilder() ContentFnBuilder {
	return createContentFnBuilder()
}

// Builder represents a selector builder
type Builder interface {
	Create() Builder
	WithGrammar(grammar grammars.Grammar) Builder
	WithToken(token Token) Builder
	WithInside(inside Inside) Builder
	WithFn(fn SelectorFn) Builder
	Now() (Selector, error)
}

// Selector represents a selector
type Selector interface {
	Grammar() grammars.Grammar
	Token() Token
	Inside() Inside
	Fn() SelectorFn
}

// SelectorFnBuilder represents the selector func builder
type SelectorFnBuilder interface {
	Create() SelectorFnBuilder
	WithSingle(single SingleSelectorFn) SelectorFnBuilder
	WithMulti(multi MultiSelectorFn) SelectorFnBuilder
	Now() (SelectorFn, error)
}

// SelectorFn represents the selector fn
type SelectorFn interface {
	IsSingle() bool
	Single() SingleSelectorFn
	IsMulti() bool
	Multi() MultiSelectorFn
}

// TokenBuilder represents a token builder
type TokenBuilder interface {
	Create() TokenBuilder
	WithName(name string) TokenBuilder
	WithReverseName(reverseName string) TokenBuilder
	WithElement(element Element) TokenBuilder
	WithContent(content uint) TokenBuilder
	Now() (Token, error)
}

// Token represents a token
type Token interface {
	Name() string
	ReverseName() string
	Element() Element
	HasContent() bool
	Content() *uint
}

// ElementBuilder represents an element builder
type ElementBuilder interface {
	Create() ElementBuilder
	WithName(name string) ElementBuilder
	WithIndex(index uint) ElementBuilder
	Now() (Element, error)
}

// Element represents an element
type Element interface {
	Name() string
	Index() uint
}

// InsideBuilder represents an inside builder
type InsideBuilder interface {
	Create() InsideBuilder
	WithFn(fn ContentFn) InsideBuilder
	WithFetchers(fetchers Fetchers) InsideBuilder
	Now() (Inside, error)
}

// Inside represents the inside
type Inside interface {
	IsFn() bool
	Fn() ContentFn
	IsFetchers() bool
	Fetchers() Fetchers
}

// FetchersBuilder represents a fetchers builder
type FetchersBuilder interface {
	Create() FetchersBuilder
	WithList(list []Fetcher) FetchersBuilder
	Now() (Fetchers, error)
}

// Fetchers represents fetchers
type Fetchers interface {
	List() []Fetcher
}

// FetcherBuilder represents a fetcher builder
type FetcherBuilder interface {
	Create() FetcherBuilder
	WithRecursive(recursive string) FetcherBuilder
	WithSelector(selector Selector) FetcherBuilder
	Now() (Fetcher, error)
}

// Fetcher represents a fetcher
type Fetcher interface {
	IsRecursive() bool
	Recursive() string
	IsSelector() bool
	Selector() Selector
}

// ContentFnBuilder represents the content func builder
type ContentFnBuilder interface {
	Create() ContentFnBuilder
	WithSingle(single SingleContentFn) ContentFnBuilder
	WithMulti(multi MultiContentFn) ContentFnBuilder
	Now() (ContentFn, error)
}

// ContentFn represents the content func
type ContentFn interface {
	IsSingle() bool
	Single() SingleContentFn
	IsMulti() bool
	Multi() MultiContentFn
}
