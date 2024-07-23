package executions

import (
	"github.com/steve-care-software/webx/engine/states/domain/hash"
	"github.com/steve-care-software/webx/engine/vms/domain/instances/layers/instructions/executions/merges"
)

type content struct {
	hash     hash.Hash
	commit   string
	rollback string
	cancel   string
	merge    merges.Merge
}

func createContentWithCommit(
	hash hash.Hash,
	commit string,
) Content {
	return createContentInternally(hash, commit, "", "", nil)
}

func createContentWithRollback(
	hash hash.Hash,
	rollback string,
) Content {
	return createContentInternally(hash, "", rollback, "", nil)
}

func createContentWithCancel(
	hash hash.Hash,
	cancel string,
) Content {
	return createContentInternally(hash, "", "", cancel, nil)
}

func createContentWithMerge(
	hash hash.Hash,
	merge merges.Merge,
) Content {
	return createContentInternally(hash, "", "", "", merge)
}

func createContentInternally(
	hash hash.Hash,
	commit string,
	rollback string,
	cancel string,
	merge merges.Merge,
) Content {
	out := content{
		hash:     hash,
		commit:   commit,
		rollback: rollback,
		cancel:   cancel,
		merge:    merge,
	}

	return &out
}

// Hash returns the hash
func (obj *content) Hash() hash.Hash {
	return obj.hash
}

// IsCommit returns true if there is a commit, false otherwise
func (obj *content) IsCommit() bool {
	return obj.commit != ""
}

// Commit returns the commit, if any
func (obj *content) Commit() string {
	return obj.commit
}

// IsRollback returns true if there is a rollback, false otherwise
func (obj *content) IsRollback() bool {
	return obj.rollback != ""
}

// Rollback returns the rollback, if any
func (obj *content) Rollback() string {
	return obj.rollback
}

// IsCancel returns true if there is a cancel, false otherwise
func (obj *content) IsCancel() bool {
	return obj.cancel != ""
}

// Cancel returns the cancel, if any
func (obj *content) Cancel() string {
	return obj.cancel
}

// IsMerge returns true if there is a merge, false otherwise
func (obj *content) IsMerge() bool {
	return obj.merge != nil
}

// Merge returns the merge, if any
func (obj *content) Merge() merges.Merge {
	return obj.merge
}
