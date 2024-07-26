package elements

import "github.com/steve-care-software/webx/engine/databases/entities/domain/hash"

type element struct {
	hash  hash.Hash
	layer hash.Hash
	bytes []byte
	str   string
}

func createElementWithLayer(
	hash hash.Hash,
	layer hash.Hash,
) Element {
	return createElementInternally(hash, layer, nil, "")
}

func createElementWithBytes(
	hash hash.Hash,
	bytes []byte,
) Element {
	return createElementInternally(hash, nil, bytes, "")
}

func createElementWithString(
	hash hash.Hash,
	str string,
) Element {
	return createElementInternally(hash, nil, nil, str)
}

func createElementInternally(
	hash hash.Hash,
	layer hash.Hash,
	bytes []byte,
	str string,
) Element {
	out := element{
		hash:  hash,
		layer: layer,
		bytes: bytes,
		str:   str,
	}

	return &out
}

// Hash returns the hash
func (obj *element) Hash() hash.Hash {
	return obj.hash
}

// IsLayer returns true if there is a layer, false otherwise
func (obj *element) IsLayer() bool {
	return obj.layer != nil
}

// Layer returns the layer, if any
func (obj *element) Layer() hash.Hash {
	return obj.layer
}

// IsBytes returns true if there is bytes, false otherwise
func (obj *element) IsBytes() bool {
	return obj.bytes != nil
}

// Bytes returns the bytes, if any
func (obj *element) Bytes() []byte {
	return obj.bytes
}

// IsString returns true if there is string, false otherwise
func (obj *element) IsString() bool {
	return obj.str != ""
}

// String returns the string, if any
func (obj *element) String() string {
	return obj.str
}
