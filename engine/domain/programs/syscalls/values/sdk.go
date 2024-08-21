package values

import (
	"github.com/steve-care-software/webx/engine/domain/programs/instructions/tokens"
	"github.com/steve-care-software/webx/engine/domain/programs/syscalls/values/parameters"
)

// Values represents values
type Values interface {
	List() []string
}

// Value represents a syscall value
type Value interface {
	IsParameter() bool
	Parameter() parameters.Parameter
	IsToken() bool
	Token() tokens.Token
}
