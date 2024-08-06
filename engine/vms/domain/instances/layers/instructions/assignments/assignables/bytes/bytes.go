package bytes

import "github.com/steve-care-software/webx/engine/cursors/domain/hash"

type bytesIns struct {
	hash      hash.Hash
	join      []string
	compare   []string
	hashBytes string
}

func createBytesWithJoin(
	hash hash.Hash,
	join []string,
) Bytes {
	return createBytesInternally(hash, join, nil, "")
}

func createBytesWithCompare(
	hash hash.Hash,
	compare []string,
) Bytes {
	return createBytesInternally(hash, nil, compare, "")
}

func createBytesWithHashBytes(
	hash hash.Hash,
	hashBytes string,
) Bytes {
	return createBytesInternally(hash, nil, nil, hashBytes)
}

func createBytesInternally(
	hash hash.Hash,
	join []string,
	compare []string,
	hashBytes string,
) Bytes {
	out := bytesIns{
		hash:      hash,
		join:      join,
		compare:   compare,
		hashBytes: hashBytes,
	}

	return &out
}

// Hash returns the hash
func (obj *bytesIns) Hash() hash.Hash {
	return obj.hash
}

// IsJoin returns true if there is a join, false otherwise
func (obj *bytesIns) IsJoin() bool {
	return obj.join != nil
}

// Join returns the join, if any
func (obj *bytesIns) Join() []string {
	return obj.join
}

// IsCompare returns true if there is a compare, false otherwise
func (obj *bytesIns) IsCompare() bool {
	return obj.compare != nil
}

// Compare returns the compare, if any
func (obj *bytesIns) Compare() []string {
	return obj.compare
}

// IsHashBytes returns true if there is a hashBytes, false otherwise
func (obj *bytesIns) IsHashBytes() bool {
	return obj.hashBytes != ""
}

// HashBytes returns the hashBytes, if any
func (obj *bytesIns) HashBytes() string {
	return obj.hashBytes
}
