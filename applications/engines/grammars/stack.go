package grammars

import "github.com/steve-care-software/syntax/domain/syntax/grammars"

type stack struct {
	token grammars.Token
	lines map[int][]byte
}
