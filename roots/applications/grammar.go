package applications

import (
	"errors"
	"fmt"

	blockchain_applications "github.com/steve-care-software/webx/blockchains/applications"
	grammar_applications "github.com/steve-care-software/webx/grammars/applications"
	contents_grammar "github.com/steve-care-software/webx/roots/domain/roots/contents/grammars"
	roots_grammar "github.com/steve-care-software/webx/roots/domain/roots/grammars"
)

type grammar struct {
	blockchainApp        blockchain_applications.Application
	blockchainAppBuilder blockchain_applications.Builder
	grammarAppBuilder    grammar_applications.Builder
	builder              roots_grammar.Builder
	contentAdapter       contents_grammar.Adapter
}

func createGrammar(
	blockchainApp blockchain_applications.Application,
	blockchainAppBuilder blockchain_applications.Builder,
	grammarAppBuilder grammar_applications.Builder,
	builder roots_grammar.Builder,
	contentAdapter contents_grammar.Adapter,
) Grammar {
	out := grammar{
		blockchainApp:        blockchainApp,
		blockchainAppBuilder: blockchainAppBuilder,
		grammarAppBuilder:    grammarAppBuilder,
		builder:              builder,
		contentAdapter:       contentAdapter,
	}

	return &out
}

// Retrieve retrieves the grammar database instance
func (app *grammar) Retrieve(context uint) (roots_grammar.Grammar, error) {
	contentKeys, err := app.blockchainApp.ContentKeys(context, GrammarDatabaseKind)
	if err != nil {
		return nil, err
	}

	list := contentKeys.List()
	if len(list) != 1 {
		str := fmt.Sprintf("%d grammar databases were expected, %d returned", 1, len(list))
		return nil, errors.New(str)
	}

	content, err := app.blockchainApp.Read(context, list[0].Content())
	if err != nil {
		return nil, err
	}

	contentGrammar, err := app.contentAdapter.ToGrammar(content)
	if err != nil {
		return nil, err
	}

	name := contentGrammar.Name()
	builder := app.builder.Create().WithName(name)
	if contentGrammar.HasHistory() {
		history := contentGrammar.History()
		builder.WithHistory(history)
	}

	return builder.Now()
}

// Application returns the grammar application
func (app *grammar) Application(context uint) (grammar_applications.Application, error) {
	grammarIns, err := app.Retrieve(context)
	if err != nil {
		return nil, err
	}

	name := grammarIns.Name()
	blockchainApp, err := app.blockchainAppBuilder.Create().WithName(name).Now()
	if err != nil {
		return nil, err
	}

	return app.grammarAppBuilder.Create().
		WithBlockchain(blockchainApp).
		Now()
}
