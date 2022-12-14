package applications

import (
	blockchain_applications "github.com/steve-care-software/webx/blockchains/applications"
	"github.com/steve-care-software/webx/programs/domain/programs/modules"
	contents_grammar "github.com/steve-care-software/webx/roots/domain/contents/grammars"
	roots_grammar "github.com/steve-care-software/webx/roots/domain/grammars"
)

type application struct {
	blockchainApp         blockchain_applications.Application
	grammar               Grammar
	program               Program
	selector              Selector
	compiler              Compiler
	grammarBuilder        roots_grammar.Builder
	contentAdapter        contents_grammar.Adapter
	contentGrammarBuilder contents_grammar.Builder
	modules               modules.Modules
}

func createApplication(
	blockchainApp blockchain_applications.Application,
	grammar Grammar,
	program Program,
	selector Selector,
	compiler Compiler,
	grammarBuilder roots_grammar.Builder,
	contentAdapter contents_grammar.Adapter,
	contentGrammarBuilder contents_grammar.Builder,
	modules modules.Modules,
) Application {
	out := application{
		blockchainApp:         blockchainApp,
		grammar:               grammar,
		program:               program,
		selector:              selector,
		compiler:              compiler,
		grammarBuilder:        grammarBuilder,
		contentAdapter:        contentAdapter,
		contentGrammarBuilder: contentGrammarBuilder,
		modules:               modules,
	}

	return &out
}

// New creates a new application
func (app *application) New(context uint, name string) error {
	grammarIns, err := app.grammarBuilder.Create().WithName(name).Now()
	if err != nil {
		return err
	}

	hash := grammarIns.Hash()
	grammarContentIns, err := app.contentGrammarBuilder.Create().WithHash(hash).WithName(name).Now()
	if err != nil {
		return err
	}

	content, err := app.contentAdapter.ToContent(grammarContentIns)
	if err != nil {
		return err
	}

	return app.blockchainApp.Write(context, hash, content, GrammarDatabaseKind)
}

// Grammar returns the grammar application
func (app *application) Grammar() Grammar {
	return app.grammar
}

// Program returns the program application
func (app *application) Program() Program {
	return app.program
}

// Selector returns the selector application
func (app *application) Selector() Selector {
	return app.selector
}

// Compiler returns the compiler application
func (app *application) Compiler() Compiler {
	return app.compiler
}
