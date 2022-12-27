package grammars

import (
	database_applications "github.com/steve-care-software/webx/databases/applications"
	"github.com/steve-care-software/webx/databases/domain/cryptography/hash"
	grammar_applications "github.com/steve-care-software/webx/grammars/applications"
	"github.com/steve-care-software/webx/grammars/domain/grammars"
	"github.com/steve-care-software/webx/grammars/domain/grammars/cardinalities"
	"github.com/steve-care-software/webx/grammars/domain/grammars/values"
	contents_grammars "github.com/steve-care-software/webx/roots/domain/grammars"
	contents_channels "github.com/steve-care-software/webx/roots/domain/grammars/channels"
	contents_elements "github.com/steve-care-software/webx/roots/domain/grammars/elements"
	contents_everythings "github.com/steve-care-software/webx/roots/domain/grammars/everythings"
	contents_tokens "github.com/steve-care-software/webx/roots/domain/grammars/tokens"
)

const (
	// KindGrammar represents the grammar kind
	KindGrammar = iota

	// KindToken represents the token kind
	KindToken

	// KindSuite represents the suite kind
	KindSuite

	// KindElement represents the element kind
	KindElement

	// KindEverything represents the everything kind
	KindEverything

	// KindChannel represents the channel kind
	KindChannel
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	grammarApp := grammar_applications.NewApplication()
	contentAdapter := contents_grammars.NewAdapter()
	contentBuilder := contents_grammars.NewBuilder()
	contentTokenAdapter := contents_tokens.NewAdapter()
	contentTokenBuilder := contents_tokens.NewBuilder()
	contentTokenLinesBuilder := contents_tokens.NewLinesBuilder()
	contentTokenLineBuilder := contents_tokens.NewLineBuilder()
	contentElementAdapter := contents_elements.NewAdapter()
	contentElementBuilder := contents_elements.NewBuilder()
	contentElementCardinalityBuilder := contents_elements.NewCardinalityBuilder()
	contentEverythingAdapter := contents_everythings.NewAdapter()
	contentEverythingBuilder := contents_everythings.NewBuilder()
	contentChannelAdapter := contents_channels.NewAdapter()
	contentChannelBuilder := contents_channels.NewBuilder()
	grammarBuilder := grammars.NewBuilder()
	grammarTokenBuilder := grammars.NewTokenBuilder()
	grammarBlockBuilder := grammars.NewBlockBuilder()
	grammarLineBuilder := grammars.NewLineBuilder()
	grammarElementBuilder := grammars.NewElementBuilder()
	grammarInstanceBuilder := grammars.NewInstanceBuilder()
	grammarExternalBuilder := grammars.NewExternalBuilder()
	grammarEverythingBuilder := grammars.NewEverythingBuilder()
	grammarChannelsBuilder := grammars.NewChannelsBuilder()
	grammarChannelBuilder := grammars.NewChannelBuilder()
	grammarChannelConditionBuilder := grammars.NewChannelConditionBuilder()
	grammarValueBuilder := values.NewBuilder()
	grammarCardinalityBuilder := cardinalities.NewBuilder()
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		grammarApp,
		contentAdapter,
		contentBuilder,
		contentTokenAdapter,
		contentTokenBuilder,
		contentTokenLinesBuilder,
		contentTokenLineBuilder,
		contentElementAdapter,
		contentElementBuilder,
		contentElementCardinalityBuilder,
		contentEverythingAdapter,
		contentEverythingBuilder,
		contentChannelAdapter,
		contentChannelBuilder,
		grammarBuilder,
		grammarTokenBuilder,
		grammarBlockBuilder,
		grammarLineBuilder,
		grammarElementBuilder,
		grammarInstanceBuilder,
		grammarExternalBuilder,
		grammarEverythingBuilder,
		grammarChannelsBuilder,
		grammarChannelBuilder,
		grammarChannelConditionBuilder,
		grammarValueBuilder,
		grammarCardinalityBuilder,
		hashAdapter,
	)
}

// Builder represents the application builder
type Builder interface {
	Create() Builder
	WithDatabase(databaseApp database_applications.Application) Builder
	Now() (Application, error)
}

// Application represents a grammar application
type Application interface {
	Retrieve(context uint, hash hash.Hash) (grammars.Grammar, error)
	Scan(context uint, suites grammars.Suites) (grammars.Grammar, error)
	ScanWithChannels(context uint, suites grammars.Suites, channels grammars.Channels) (grammars.Grammar, error)
	Insert(context uint, grammar grammars.Grammar) error
}
