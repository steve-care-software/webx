package outputs

import "github.com/steve-care-software/syntax/domain/syntax/programs"

type output struct {
	program   programs.Program
	remaining Remaining
}

func createOutput(
	program programs.Program,
) Output {
	return createOutputInternally(program, nil)
}

func createOutputWithRemaining(
	program programs.Program,
	remaining Remaining,
) Output {
	return createOutputInternally(program, remaining)
}

func createOutputInternally(
	program programs.Program,
	remaining Remaining,
) Output {
	out := output{
		program:   program,
		remaining: remaining,
	}

	return &out
}

// Program returns the program
func (obj *output) Program() programs.Program {
	return obj.program
}

// HasRemaining returns true if there is a remaining, false otherwise
func (obj *output) HasRemaining() bool {
	return obj.remaining != nil
}

// Remaining returns the remaining, if any
func (obj *output) Remaining() Remaining {
	return obj.remaining
}
