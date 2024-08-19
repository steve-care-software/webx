package blocks

import (
	"github.com/steve-care-software/webx/engine/domain/grammars/blocks/lines"
	"github.com/steve-care-software/webx/engine/domain/grammars/blocks/suites"
)

type block struct {
	name   string
	lines  lines.Lines
	suites suites.Suites
}

func createBlock(
	name string,
	lines lines.Lines,
) Block {
	return createBlockInternally(name, lines, nil)
}

func createBlockWithSuites(
	name string,
	lines lines.Lines,
	suites suites.Suites,
) Block {
	return createBlockInternally(name, lines, suites)
}

func createBlockInternally(
	name string,
	lines lines.Lines,
	suites suites.Suites,
) Block {
	out := block{
		name:   name,
		lines:  lines,
		suites: suites,
	}

	return &out
}

// Name returns the name
func (obj *block) Name() string {
	return obj.name
}

// Lines returns the lines
func (obj *block) Lines() lines.Lines {
	return obj.lines
}

// HasSuites returns true if there is a list, false otherwise
func (obj *block) HasSuites() bool {
	return obj.suites != nil
}

// Suites returns the suites, if any
func (obj *block) Suites() suites.Suites {
	return obj.suites
}
