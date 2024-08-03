package files

import "github.com/steve-care-software/webx/engine/databases/hashes/domain/hash"

type file struct {
	hash   hash.Hash
	close  string
	delete string
}

func createFileWithClose(
	hash hash.Hash,
	close string,
) File {
	return createFileInternally(hash, close, "")
}

func createFileWithDelete(
	hash hash.Hash,
	delete string,
) File {
	return createFileInternally(hash, "", delete)
}

func createFileInternally(
	hash hash.Hash,
	close string,
	delete string,
) File {
	out := file{
		hash:   hash,
		close:  close,
		delete: delete,
	}

	return &out
}

// Hash represents the hash
func (obj *file) Hash() hash.Hash {
	return obj.hash
}

// IsClose returns true if close, false otherwise
func (obj *file) IsClose() bool {
	return obj.close != ""
}

// Close returns the close, if any
func (obj *file) Close() string {
	return obj.close
}

// IsDelete returns true if delete, false otherwise
func (obj *file) IsDelete() bool {
	return obj.delete != ""
}

// Delete returns the delete, if any
func (obj *file) Delete() string {
	return obj.delete
}
