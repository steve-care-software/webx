package applications

import (
	"errors"
	"fmt"

	database_applications "github.com/steve-care-software/webx/databases/applications"
	grammar_applications "github.com/steve-care-software/webx/roots/applications/grammars"
	contents_grammar "github.com/steve-care-software/webx/roots/domain/roots/contents/grammars"
	roots_grammar "github.com/steve-care-software/webx/roots/domain/roots/grammars"
)

type grammar struct {
	databaseApp        database_applications.Application
	databaseAppBuilder database_applications.Builder
	grammarAppBuilder  grammar_applications.Builder
	builder            roots_grammar.Builder
	contentAdapter     contents_grammar.Adapter
}

func createGrammar(
	databaseApp database_applications.Application,
	databaseAppBuilder database_applications.Builder,
	grammarAppBuilder grammar_applications.Builder,
	builder roots_grammar.Builder,
	contentAdapter contents_grammar.Adapter,
) Grammar {
	out := grammar{
		databaseApp:        databaseApp,
		databaseAppBuilder: databaseAppBuilder,
		grammarAppBuilder:  grammarAppBuilder,
		builder:            builder,
		contentAdapter:     contentAdapter,
	}

	return &out
}

// Retrieve retrieves the grammar database instance
func (app *grammar) Retrieve(context uint) (roots_grammar.Grammar, error) {
	contentKeys, err := app.databaseApp.ContentKeys(context, GrammarDatabaseKind)
	if err != nil {
		return nil, err
	}

	list := contentKeys.List()
	if len(list) != 1 {
		str := fmt.Sprintf("%d grammar databases were expected, %d returned", 1, len(list))
		return nil, errors.New(str)
	}

	content, err := app.databaseApp.Read(context, list[0].Content())
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
	databaseApp, err := app.databaseAppBuilder.Create().WithName(name).Now()
	if err != nil {
		return nil, err
	}

	return app.grammarAppBuilder.Create().
		WithBlockchain(databaseApp).
		Now()
}
