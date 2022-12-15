package grammars

import "github.com/steve-care-software/webx/roots/domain/grammars/grammars"

type stack struct {
	token grammars.Token
	lines map[int][]byte
}
