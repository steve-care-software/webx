package grammars

import (
	"errors"

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

type builder struct {
	grammarApp                       grammar_applications.Application
	contentAdapter                   contents_grammars.Adapter
	contentBuilder                   contents_grammars.Builder
	contentTokenAdapter              contents_tokens.Adapter
	contentTokenBuilder              contents_tokens.Builder
	contentTokenLinesBuilder         contents_tokens.LinesBuilder
	contentTokenLineBuilder          contents_tokens.LineBuilder
	contentElementAdapter            contents_elements.Adapter
	contentElementBuilder            contents_elements.Builder
	contentElementCardinalityBuilder contents_elements.CardinalityBuilder
	contentEverythingAdapter         contents_everythings.Adapter
	contentEverythingBuilder         contents_everythings.Builder
	contentChannelAdapter            contents_channels.Adapter
	contentChannelBuilder            contents_channels.Builder
	builder                          grammars.Builder
	grammarTokenBuilder              grammars.TokenBuilder
	grammarBlockBuilder              grammars.BlockBuilder
	grammarLineBuilder               grammars.LineBuilder
	grammarElementBuilder            grammars.ElementBuilder
	grammarInstanceBuilder           grammars.InstanceBuilder
	grammarExternalBuilder           grammars.ExternalBuilder
	grammarEverythingBuilder         grammars.EverythingBuilder
	grammarChannelsBuilder           grammars.ChannelsBuilder
	grammarChannelBuilder            grammars.ChannelBuilder
	grammarChannelConditionBuilder   grammars.ChannelConditionBuilder
	grammarValueBuilder              values.Builder
	grammarCardinalityBuilder        cardinalities.Builder
	hashAdapter                      hash.Adapter
	databaseApp                      database_applications.Application
}

func createBuilder(
	grammarApp grammar_applications.Application,
	contentAdapter contents_grammars.Adapter,
	contentBuilder contents_grammars.Builder,
	contentTokenAdapter contents_tokens.Adapter,
	contentTokenBuilder contents_tokens.Builder,
	contentTokenLinesBuilder contents_tokens.LinesBuilder,
	contentTokenLineBuilder contents_tokens.LineBuilder,
	contentElementAdapter contents_elements.Adapter,
	contentElementBuilder contents_elements.Builder,
	contentElementCardinalityBuilder contents_elements.CardinalityBuilder,
	contentEverythingAdapter contents_everythings.Adapter,
	contentEverythingBuilder contents_everythings.Builder,
	contentChannelAdapter contents_channels.Adapter,
	contentChannelBuilder contents_channels.Builder,
	grammarBuilder grammars.Builder,
	grammarTokenBuilder grammars.TokenBuilder,
	grammarBlockBuilder grammars.BlockBuilder,
	grammarLineBuilder grammars.LineBuilder,
	grammarElementBuilder grammars.ElementBuilder,
	grammarInstanceBuilder grammars.InstanceBuilder,
	grammarExternalBuilder grammars.ExternalBuilder,
	grammarEverythingBuilder grammars.EverythingBuilder,
	grammarChannelsBuilder grammars.ChannelsBuilder,
	grammarChannelBuilder grammars.ChannelBuilder,
	grammarChannelConditionBuilder grammars.ChannelConditionBuilder,
	grammarValueBuilder values.Builder,
	grammarCardinalityBuilder cardinalities.Builder,
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		grammarApp:                       grammarApp,
		contentAdapter:                   contentAdapter,
		contentBuilder:                   contentBuilder,
		contentTokenAdapter:              contentTokenAdapter,
		contentTokenBuilder:              contentTokenBuilder,
		contentTokenLinesBuilder:         contentTokenLinesBuilder,
		contentTokenLineBuilder:          contentTokenLineBuilder,
		contentElementAdapter:            contentElementAdapter,
		contentElementBuilder:            contentElementBuilder,
		contentElementCardinalityBuilder: contentElementCardinalityBuilder,
		contentEverythingAdapter:         contentEverythingAdapter,
		contentEverythingBuilder:         contentEverythingBuilder,
		contentChannelAdapter:            contentChannelAdapter,
		contentChannelBuilder:            contentChannelBuilder,
		builder:                          grammarBuilder,
		grammarTokenBuilder:              grammarTokenBuilder,
		grammarBlockBuilder:              grammarBlockBuilder,
		grammarLineBuilder:               grammarLineBuilder,
		grammarElementBuilder:            grammarElementBuilder,
		grammarInstanceBuilder:           grammarInstanceBuilder,
		grammarExternalBuilder:           grammarExternalBuilder,
		grammarEverythingBuilder:         grammarEverythingBuilder,
		grammarChannelsBuilder:           grammarChannelsBuilder,
		grammarChannelBuilder:            grammarChannelBuilder,
		grammarChannelConditionBuilder:   grammarChannelConditionBuilder,
		grammarValueBuilder:              grammarValueBuilder,
		grammarCardinalityBuilder:        grammarCardinalityBuilder,
		hashAdapter:                      hashAdapter,
		databaseApp:                      nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.grammarApp,
		app.contentAdapter,
		app.contentBuilder,
		app.contentTokenAdapter,
		app.contentTokenBuilder,
		app.contentTokenLinesBuilder,
		app.contentTokenLineBuilder,
		app.contentElementAdapter,
		app.contentElementBuilder,
		app.contentElementCardinalityBuilder,
		app.contentEverythingAdapter,
		app.contentEverythingBuilder,
		app.contentChannelAdapter,
		app.contentChannelBuilder,
		app.builder,
		app.grammarTokenBuilder,
		app.grammarBlockBuilder,
		app.grammarLineBuilder,
		app.grammarElementBuilder,
		app.grammarInstanceBuilder,
		app.grammarExternalBuilder,
		app.grammarEverythingBuilder,
		app.grammarChannelsBuilder,
		app.grammarChannelBuilder,
		app.grammarChannelConditionBuilder,
		app.grammarValueBuilder,
		app.grammarCardinalityBuilder,
		app.hashAdapter,
	)
}

// WithDatabase adds a database to the builder
func (app *builder) WithDatabase(databaseApp database_applications.Application) Builder {
	app.databaseApp = databaseApp
	return app
}

// Now builds a new Application instance
func (app *builder) Now() (Application, error) {
	if app.databaseApp == nil {
		return nil, errors.New("the database application is mandatory in order to build an Application instance")
	}

	return createApplication(
		app.grammarApp,
		app.databaseApp,
		app.contentAdapter,
		app.contentBuilder,
		app.contentTokenAdapter,
		app.contentTokenBuilder,
		app.contentTokenLinesBuilder,
		app.contentTokenLineBuilder,
		app.contentElementAdapter,
		app.contentElementBuilder,
		app.contentElementCardinalityBuilder,
		app.contentEverythingAdapter,
		app.contentEverythingBuilder,
		app.contentChannelAdapter,
		app.contentChannelBuilder,
		app.builder,
		app.grammarTokenBuilder,
		app.grammarBlockBuilder,
		app.grammarLineBuilder,
		app.grammarElementBuilder,
		app.grammarInstanceBuilder,
		app.grammarExternalBuilder,
		app.grammarEverythingBuilder,
		app.grammarChannelsBuilder,
		app.grammarChannelBuilder,
		app.grammarChannelConditionBuilder,
		app.grammarValueBuilder,
		app.grammarCardinalityBuilder,
		app.hashAdapter,
	), nil
}
