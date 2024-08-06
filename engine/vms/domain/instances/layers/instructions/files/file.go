package files

import "github.com/steve-care-software/webx/engine/cursors/domain/hash"

type file struct {
	hash  hash.Hash
	close string
}

func createFileWithClose(
	hash hash.Hash,
	close string,
) File {
	return createFileInternally(hash, close)
}

func createFileInternally(
	hash hash.Hash,
	close string,
) File {
	out := file{
		hash:  hash,
		close: close,
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
