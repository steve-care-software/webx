package applications

import "github.com/steve-care-software/historydb/domain/hash"

type context struct {
	dbPath     []string
	executions []hash.Hash
}
