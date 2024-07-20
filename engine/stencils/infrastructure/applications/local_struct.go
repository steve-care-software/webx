package applications

import "github.com/steve-care-software/webx/engine/states/domain/hash"

type context struct {
	dbPath     []string
	executions []hash.Hash
}
