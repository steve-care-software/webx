package grammars

import (
	"github.com/steve-care-software/webx/engine/domain/programs/grammars/blocks"
	"github.com/steve-care-software/webx/engine/domain/programs/grammars/blocks/lines/tokens/elements"
	"github.com/steve-care-software/webx/engine/domain/programs/grammars/resources"
	"github.com/steve-care-software/webx/engine/domain/programs/grammars/rules"
	"github.com/steve-care-software/webx/engine/domain/programs/grammars/spacers"
)

type grammar struct {
	version   uint
	root      elements.Element
	rules     rules.Rules
	blocks    blocks.Blocks
	omissions elements.Elements
	resources resources.Resources
	spacers   spacers.Spacers
}

func createGrammar(
	version uint,
	root elements.Element,
	rules rules.Rules,
	blocks blocks.Blocks,
) Grammar {
	return createGrammarInternally(version, root, rules, blocks, nil, nil, nil)
}

func createGrammarWithOmissions(
	version uint,
	root elements.Element,
	rules rules.Rules,
	blocks blocks.Blocks,
	omissions elements.Elements,
) Grammar {
	return createGrammarInternally(version, root, rules, blocks, omissions, nil, nil)
}

func createGrammarWithResources(
	version uint,
	root elements.Element,
	rules rules.Rules,
	blocks blocks.Blocks,
	resources resources.Resources,
) Grammar {
	return createGrammarInternally(version, root, rules, blocks, nil, resources, nil)
}

func createGrammarWithSpacers(
	version uint,
	root elements.Element,
	rules rules.Rules,
	blocks blocks.Blocks,
	spacers spacers.Spacers,
) Grammar {
	return createGrammarInternally(version, root, rules, blocks, nil, nil, spacers)
}

func createGrammarWithOmissionsAndSpacers(
	version uint,
	root elements.Element,
	rules rules.Rules,
	blocks blocks.Blocks,
	omissions elements.Elements,
	spacers spacers.Spacers,
) Grammar {
	return createGrammarInternally(version, root, rules, blocks, omissions, nil, spacers)
}

func createGrammarWithResourcesAndSpacers(
	version uint,
	root elements.Element,
	rules rules.Rules,
	blocks blocks.Blocks,
	resources resources.Resources,
	spacers spacers.Spacers,
) Grammar {
	return createGrammarInternally(version, root, rules, blocks, nil, resources, spacers)
}

func createGrammarWithOmissionsAndResources(
	version uint,
	root elements.Element,
	rules rules.Rules,
	blocks blocks.Blocks,
	omissions elements.Elements,
	resources resources.Resources,
) Grammar {
	return createGrammarInternally(version, root, rules, blocks, omissions, resources, nil)
}

func createGrammarWithOmissionsAndResourcesAndSpacers(
	version uint,
	root elements.Element,
	rules rules.Rules,
	blocks blocks.Blocks,
	omissions elements.Elements,
	resources resources.Resources,
	spacers spacers.Spacers,
) Grammar {
	return createGrammarInternally(version, root, rules, blocks, omissions, resources, spacers)
}

func createGrammarInternally(
	version uint,
	root elements.Element,
	rules rules.Rules,
	blocks blocks.Blocks,
	omissions elements.Elements,
	resources resources.Resources,
	spacers spacers.Spacers,
) Grammar {
	out := grammar{
		version:   version,
		root:      root,
		rules:     rules,
		blocks:    blocks,
		omissions: omissions,
		resources: resources,
		spacers:   spacers,
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

// HasResources returns true if there is resources, false otherwise
func (obj *grammar) HasResources() bool {
	return obj.resources != nil
}

// Resources returns the resources, if any
func (obj *grammar) Resources() resources.Resources {
	return obj.resources
}

// HasSpacers returns true if there is spacers, false otherwise
func (obj *grammar) HasSpacers() bool {
	return obj.spacers != nil
}

// Spacers returns the spacers, if any
func (obj *grammar) Spacers() spacers.Spacers {
	return obj.spacers
}
