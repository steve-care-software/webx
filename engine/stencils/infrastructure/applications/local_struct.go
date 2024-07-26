package applications

import "github.com/steve-care-software/webx/engine/databases/entities/domain/hash"

type context struct {
	dbPath     []string
	executions []hash.Hash
}
