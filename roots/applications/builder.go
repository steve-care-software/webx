package applications

import (
	"errors"

	database_applications "github.com/steve-care-software/webx/databases/applications"
	grammar_applications "github.com/steve-care-software/webx/roots/applications/grammars"
	"github.com/steve-care-software/webx/roots/domain/programs/programs/modules"
	contents_grammar "github.com/steve-care-software/webx/roots/domain/roots/contents/grammars"
	roots_grammar "github.com/steve-care-software/webx/roots/domain/roots/grammars"
)

type builder struct {
	databaseAppBuilder    database_applications.Builder
	grammarAppBuilder     grammar_applications.Builder
	grammarBuilder        roots_grammar.Builder
	contentGrammarAdapter contents_grammar.Adapter
	contentGrammarBuilder contents_grammar.Builder
	blockchain            database_applications.Application
	modules               modules.Modules
}

func createBuilder(
	databaseAppBuilder database_applications.Builder,
	grammarAppBuilder grammar_applications.Builder,
	grammarBuilder roots_grammar.Builder,
	contentGrammarAdapter contents_grammar.Adapter,
	contentGrammarBuilder contents_grammar.Builder,
) Builder {
	out := builder{
		databaseAppBuilder:    databaseAppBuilder,
		grammarAppBuilder:     grammarAppBuilder,
		grammarBuilder:        grammarBuilder,
		contentGrammarAdapter: contentGrammarAdapter,
		contentGrammarBuilder: contentGrammarBuilder,
		blockchain:            nil,
		modules:               nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.databaseAppBuilder,
		app.grammarAppBuilder,
		app.grammarBuilder,
		app.contentGrammarAdapter,
		app.contentGrammarBuilder,
	)
}

// WithBlockchain adds a blockchain to the builder
func (app *builder) WithBlockchain(blockchain database_applications.Application) Builder {
	app.blockchain = blockchain
	return app
}

// WithModules add modules to the builder
func (app *builder) WithModules(modules modules.Modules) Builder {
	app.modules = modules
	return app
}

// Now builds a new Application instance
func (app *builder) Now() (Application, error) {
	if app.blockchain == nil {
		return nil, errors.New("the blockchain application is mandatory in order to build an Application instance")
	}

	if app.modules == nil {
		return nil, errors.New("the modules are mandatory in order to build an Application instance")
	}

	grammar := createGrammar(
		app.blockchain,
		app.databaseAppBuilder,
		app.grammarAppBuilder,
		app.grammarBuilder,
		app.contentGrammarAdapter,
	)

	return createApplication(
		app.blockchain,
		grammar,
		nil,
		nil,
		nil,
		app.grammarBuilder,
		app.contentGrammarAdapter,
		app.contentGrammarBuilder,
		app.modules,
	), nil
}
