package reads

import "github.com/steve-care-software/webx/engine/states/domain/hash"

type read struct {
	hash       hash.Hash
	identifier string
	index      string
	length     string
}

func createRead(
	hash hash.Hash,
	identifier string,
) Read {
	return createReadInternally(hash, identifier, "", "")
}

func createReadWithIndex(
	hash hash.Hash,
	identifier string,
	index string,
) Read {
	return createReadInternally(hash, identifier, index, "")
}

func createReadWithLength(
	hash hash.Hash,
	identifier string,
	length string,
) Read {
	return createReadInternally(hash, identifier, "", length)
}

func createReadWithIndexAndLength(
	hash hash.Hash,
	identifier string,
	index string,
	length string,
) Read {
	return createReadInternally(hash, identifier, index, length)
}

func createReadInternally(
	hash hash.Hash,
	identifier string,
	index string,
	length string,
) Read {
	out := read{
		hash:       hash,
		identifier: identifier,
		index:      index,
		length:     length,
	}

	return &out
}

// Hash returns the hash
func (obj *read) Hash() hash.Hash {
	return obj.hash
}

// Identifier returns the identifier
func (obj *read) Identifier() string {
	return obj.identifier
}

// HasIndex returns true if there is an index, false otherwise
func (obj *read) HasIndex() bool {
	return obj.index != ""
}

// Index returns the index, if any
func (obj *read) Index() string {
	return obj.index
}

// HasLength returns true if there is a length, false otherwise
func (obj *read) HasLength() bool {
	return obj.length != ""
}

// Length returns the length, if any
func (obj *read) Length() string {
	return obj.length
}
