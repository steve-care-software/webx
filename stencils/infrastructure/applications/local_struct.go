package applications

import "github.com/steve-care-software/datastencil/states/domain/hash"

type context struct {
	dbPath     []string
	executions []hash.Hash
}
