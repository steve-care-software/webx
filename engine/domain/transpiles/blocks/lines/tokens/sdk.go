package tokens

import (
	"github.com/steve-care-software/webx/engine/domain/transpiles/blocks/lines/tokens/pointers"
	"github.com/steve-care-software/webx/engine/domain/transpiles/blocks/lines/tokens/updates"
)

// Tokens represents tokens
type Tokens interface {
	List() []Token
}

// Token represents a token
type Token interface {
	IsUpdate() bool
	Update() updates.Update
	IsDelete() bool
	Delete() pointers.Pointer
	IsInsert() bool
	Insert() pointers.Pointer
}
