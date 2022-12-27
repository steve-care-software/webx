package instances

import (
	"github.com/steve-care-software/webx/grammars/domain/grammars"
	"github.com/steve-care-software/webx/grammars/domain/grammars/cardinalities"
	grammar_values "github.com/steve-care-software/webx/grammars/domain/grammars/values"
)

// NewGrammar creates a new gramar instance
func NewGrammar() grammars.Grammar {
	grammarIns := createGrammar(
		grammars.NewBuilder(),
		grammars.NewChannelsBuilder(),
		grammars.NewChannelBuilder(),
		grammars.NewInstanceBuilder(),
		grammars.NewEverythingBuilder(),
		grammars.NewTokensBuilder(),
		grammars.NewTokenBuilder(),
		grammars.NewSuitesBuilder(),
		grammars.NewSuiteBuilder(),
		grammars.NewBlockBuilder(),
		grammars.NewLineBuilder(),
		grammars.NewElementBuilder(),
		grammar_values.NewBuilder(),
		cardinalities.NewBuilder(),
	)

	ins, err := grammarIns.Execute()
	if err != nil {
		panic(err)
	}

	return ins
}
