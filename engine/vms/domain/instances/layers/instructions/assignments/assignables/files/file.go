package files

import (
	"github.com/steve-care-software/webx/engine/databases/hashes/domain/hash"
	"github.com/steve-care-software/webx/engine/vms/domain/instances/layers/instructions/assignments/assignables/files/opens"
	"github.com/steve-care-software/webx/engine/vms/domain/instances/layers/instructions/assignments/assignables/files/reads"
)

type file struct {
	hash   hash.Hash
	open   opens.Open
	read   reads.Read
	exists string
	length string
}

func createFileWithOpen(
	hash hash.Hash,
	open opens.Open,
) File {
	return createFileInternally(hash, open, nil, "", "")
}

func createFileWithRead(
	hash hash.Hash,
	read reads.Read,
) File {
	return createFileInternally(hash, nil, read, "", "")
}

func createFileWithExists(
	hash hash.Hash,
	exists string,
) File {
	return createFileInternally(hash, nil, nil, exists, "")
}

func createFileWithLength(
	hash hash.Hash,
	length string,
) File {
	return createFileInternally(hash, nil, nil, "", length)
}

func createFileInternally(
	hash hash.Hash,
	open opens.Open,
	read reads.Read,
	exists string,
	length string,
) File {
	out := file{
		hash:   hash,
		open:   open,
		read:   read,
		exists: exists,
		length: length,
	}

	return &out
}

// Hash returns the hash
func (obj *file) Hash() hash.Hash {
	return obj.hash
}

// IsOpen returns true if open, false otherwise
func (obj *file) IsOpen() bool {
	return obj.open != nil
}

// Open returns the open, if any
func (obj *file) Open() opens.Open {
	return obj.open
}

// IsRead returns true if read, false otherwise
func (obj *file) IsRead() bool {
	return obj.read != nil
}

// Read returns the read, if any
func (obj *file) Read() reads.Read {
	return obj.read
}

// IsExists returns true if exists, false otherwise
func (obj *file) IsExists() bool {
	return obj.exists != ""
}

// Exists returns the exists, if any
func (obj *file) Exists() string {
	return obj.exists
}

// IsLength returns true if length, false otherwise
func (obj *file) IsLength() bool {
	return obj.length != ""
}

// Length returns the length, if any
func (obj *file) Length() string {
	return obj.length
}
