package applications

import "github.com/steve-care-software/webx/grammars/domain/grammars"

type stack struct {
	token grammars.Token
	lines map[int][]byte
}
