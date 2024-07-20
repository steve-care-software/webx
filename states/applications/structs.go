package applications

import (
	"github.com/steve-care-software/datastencil/states/domain/databases/commits"
	"github.com/steve-care-software/datastencil/states/domain/databases/commits/executions"
	"github.com/steve-care-software/datastencil/states/domain/databases/metadatas"
)

type contexts struct {
	path       []string
	executions []executionData
	metaData   metadatas.MetaData
}

type executionData struct {
	execution executions.Execution
	bytes     []byte
}

type commit struct {
	path     []string
	commits  []commits.Commit
	metaData metadatas.MetaData
}
