package commands

import "github.com/steve-care-software/syntax/domain/bytes/grammars"

type command struct {
	grammar grammars.Grammar
	content Content
}

func createCommand(
	grammar grammars.Grammar,
	content Content,
) Command {
	out := command{
		grammar: grammar,
		content: content,
	}

	return &out
}

// Grammar returns the grammar
func (obj *command) Grammar() grammars.Grammar {
	return obj.grammar
}

// Content returns the content
func (obj *command) Content() Content {
	return obj.content
}
