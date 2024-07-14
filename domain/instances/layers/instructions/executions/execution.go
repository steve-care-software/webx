package executions

import (
	"github.com/steve-care-software/datastencil/domain/instances/layers/instructions/executions/merges"
	"github.com/steve-care-software/historydb/domain/databases/commits"
	"github.com/steve-care-software/historydb/domain/hash"
)

type execution struct {
	hash     hash.Hash
	commit   commits.Commit
	rollback string
	cancel   string
	merge    merges.Merge
}

func createExecutionWithCommit(
	hash hash.Hash,
	commit commits.Commit,
) Execution {
	return createExecutionInternally(hash, commit, "", "", nil)
}

func createExecutionWithRollback(
	hash hash.Hash,
	rollback string,
) Execution {
	return createExecutionInternally(hash, nil, rollback, "", nil)
}

func createExecutionWithCancel(
	hash hash.Hash,
	cancel string,
) Execution {
	return createExecutionInternally(hash, nil, "", cancel, nil)
}

func createExecutionWithMerge(
	hash hash.Hash,
	merge merges.Merge,
) Execution {
	return createExecutionInternally(hash, nil, "", "", merge)
}

func createExecutionInternally(
	hash hash.Hash,
	commit commits.Commit,
	rollback string,
	cancel string,
	merge merges.Merge,
) Execution {
	out := execution{
		hash:     hash,
		commit:   commit,
		rollback: rollback,
		cancel:   cancel,
		merge:    merge,
	}

	return &out
}

// Hash returns the hash
func (obj *execution) Hash() hash.Hash {
	return obj.hash
}

// IsCommit returns true if there is a commit, false otherwise
func (obj *execution) IsCommit() bool {
	return obj.commit != nil
}

// Commit returns the commit, if any
func (obj *execution) Commit() commits.Commit {
	return obj.commit
}

// IsRollback returns true if there is a rollback, false otherwise
func (obj *execution) IsRollback() bool {
	return obj.rollback != ""
}

// Rollback returns the rollback, if any
func (obj *execution) Rollback() string {
	return obj.rollback
}

// IsCancel returns true if there is a cancel, false otherwise
func (obj *execution) IsCancel() bool {
	return obj.cancel != ""
}

// Cancel returns the cancel, if any
func (obj *execution) Cancel() string {
	return obj.cancel
}

// IsMerge returns true if there is a merge, false otherwise
func (obj *execution) IsMerge() bool {
	return obj.merge != nil
}

// Merge returns the merge, if any
func (obj *execution) Merge() merges.Merge {
	return obj.merge
}
