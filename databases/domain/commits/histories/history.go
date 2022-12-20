package histories

import "github.com/steve-care-software/webx/databases/domain/cryptography/hash"

type history struct {
	commit hash.Hash
	score  uint
}

func createHistory(
	commit hash.Hash,
	score uint,
) History {
	out := history{
		commit: commit,
		score:  score,
	}

	return &out
}

// Commit returns the commit
func (obj *history) Commit() hash.Hash {
	return obj.commit
}

// Score returns the score
func (obj *history) Score() uint {
	return obj.score
}
