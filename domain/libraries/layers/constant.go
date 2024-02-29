package layers

import "github.com/steve-care-software/datastencil/domain/hash"

type constant struct {
	hash      hash.Hash
	boolValue *bool
	data      []byte
}

func createConstantWithBool(
	hash hash.Hash,
	boolValue *bool,
) Constant {
	return createConstantInternally(hash, boolValue, nil)
}

func createConstantWithData(
	hash hash.Hash,
	data []byte,
) Constant {
	return createConstantInternally(hash, nil, data)
}

func createConstantInternally(
	hash hash.Hash,
	boolValue *bool,
	data []byte,
) Constant {
	out := constant{
		hash:      hash,
		boolValue: boolValue,
		data:      data,
	}

	return &out
}

// Hash returns the hash
func (obj *constant) Hash() hash.Hash {
	return obj.hash
}

// IsBool returns true if there is a bool, false otherwise
func (obj *constant) IsBool() bool {
	return obj.boolValue != nil
}

// Bool returns the bool, if any
func (obj *constant) Bool() *bool {
	return obj.boolValue
}

// IsBytes returns true if there is bytes, false otherwise
func (obj *constant) IsBytes() bool {
	return obj.data != nil
}

// Bytes returns the bytes, if any
func (obj *constant) Bytes() []byte {
	return obj.data
}
