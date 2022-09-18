package programs

import (
	"github.com/steve-care-software/logics/domain/programs/instructions"
)

// Program represents a program
type Program interface {
	Compose() string
	Instructions() instructions.Instructions
	HasInput() bool
	Input() []string
	HasOutput() bool
	Output() []string
}
