package blocks

import (
	"github.com/steve-care-software/webx/engine/domain/programs/grammars/blocks/lines"
	"github.com/steve-care-software/webx/engine/domain/programs/grammars/blocks/suites"
)

type block struct {
	name   string
	line   lines.Line
	lines  lines.Lines
	suites suites.Suites
}

func createBlockWithLine(
	name string,
	line lines.Line,
) Block {
	return createBlockInternally(name, line, nil, nil)
}

func createBlockWithLineAndSuites(
	name string,
	line lines.Line,
	suites suites.Suites,
) Block {
	return createBlockInternally(name, line, nil, suites)
}

func createBlockWithLines(
	name string,
	lines lines.Lines,
) Block {
	return createBlockInternally(name, nil, lines, nil)
}

func createBlockWithLinesAndSuites(
	name string,
	lines lines.Lines,
	suites suites.Suites,
) Block {
	return createBlockInternally(name, nil, lines, suites)
}

func createBlockInternally(
	name string,
	line lines.Line,
	lines lines.Lines,
	suites suites.Suites,
) Block {
	out := block{
		name:   name,
		line:   line,
		lines:  lines,
		suites: suites,
	}

	return &out
}

// Name returns the name
func (obj *block) Name() string {
	return obj.name
}

// HasLine returns true if there is line, false otherwise
func (obj *block) HasLine() bool {
	return obj.line != nil
}

// Line returns the line, if any
func (obj *block) Line() lines.Line {
	return obj.line
}

// HasLines returns true if there is lines, false otherwise
func (obj *block) HasLines() bool {
	return obj.lines != nil
}

// Lines returns the lines
func (obj *block) Lines() lines.Lines {
	return obj.lines
}

// HasSuites returns true if there is suites, false otherwise
func (obj *block) HasSuites() bool {
	return obj.suites != nil
}

// Suites returns the suites, if any
func (obj *block) Suites() suites.Suites {
	return obj.suites
}
