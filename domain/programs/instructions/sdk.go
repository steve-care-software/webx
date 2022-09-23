package instructions

import (
	"github.com/steve-care-software/syntax/domain/programs/instructions/applications"
)

// Instructions represents instructions
type Instructions interface {
	List() []Instruction
	Compose() string
}

// Instruction represents an instruction
type Instruction interface {
	Content() Content
	Compose() string
}

// Content represents an instruction content
type Content interface {
	IsAssignment() bool
	Assignment() applications.Assignment
	IsExecution() bool
	Execution() applications.Application
	IsDelete() bool
	Delete() applications.Application
	IsSetPath() bool
	SetPath() string
}
