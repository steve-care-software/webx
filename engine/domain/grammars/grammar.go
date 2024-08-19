package grammars

import (
	"github.com/steve-care-software/webx/engine/domain/grammars/blocks"
	"github.com/steve-care-software/webx/engine/domain/grammars/blocks/lines/tokens/elements"
	"github.com/steve-care-software/webx/engine/domain/grammars/rules"
)

type grammar struct {
	version   uint
	root      elements.Element
	rules     rules.Rules
	blocks    blocks.Blocks
	omissions elements.Elements
}

func createGrammar(
	version uint,
	root elements.Element,
	rules rules.Rules,
	blocks blocks.Blocks,
) Grammar {
	return createGrammarInternally(version, root, rules, blocks, nil)
}

func createGrammarWithOmissions(
	version uint,
	root elements.Element,
	rules rules.Rules,
	blocks blocks.Blocks,
	omissions elements.Elements,
) Grammar {
	return createGrammarInternally(version, root, rules, blocks, omissions)
}

func createGrammarInternally(
	version uint,
	root elements.Element,
	rules rules.Rules,
	blocks blocks.Blocks,
	omissions elements.Elements,
) Grammar {
	out := grammar{
		version:   version,
		root:      root,
		rules:     rules,
		blocks:    blocks,
		omissions: omissions,
	}

	return &out
}

// Version returns the version
func (obj *grammar) Version() uint {
	return obj.version
}

// Root returns the root
func (obj *grammar) Root() elements.Element {
	return obj.root
}

// Rules returns the rules
func (obj *grammar) Rules() rules.Rules {
	return obj.rules
}

// Blocks returns the blocks
func (obj *grammar) Blocks() blocks.Blocks {
	return obj.blocks
}

// HasOmissions returns true if there is omissions, false otherwise
func (obj *grammar) HasOmissions() bool {
	return obj.omissions != nil
}

// Omissions returns the omissions, if any
func (obj *grammar) Omissions() elements.Elements {
	return obj.omissions
}
