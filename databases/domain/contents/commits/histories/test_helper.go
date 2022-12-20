package histories

import (
	"fmt"

	"github.com/steve-care-software/webx/databases/domain/cryptography/hash"
)

// NewHistoryForTests creates a new history for tests
func NewHistoryForTests(score uint) History {
	pHash, err := hash.NewAdapter().FromBytes([]byte(fmt.Sprintf("this is a commit hash: %d", score)))
	if err != nil {
		panic(err)
	}

	ins, err := NewBuilder().WithCommit(*pHash).WithScore(score).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
