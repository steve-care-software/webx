package grammars

import "github.com/steve-care-software/webx/domain/grammars"

type stack struct {
	token grammars.Token
	lines map[int][]byte
}
