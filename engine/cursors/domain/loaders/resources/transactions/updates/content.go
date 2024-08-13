package updates

import (
	"github.com/steve-care-software/webx/engine/cursors/domain/hash"
)

type content struct {
	hash     hash.Hash
	name     string
	data     []byte
	addition []hash.Hash
	removal  []hash.Hash
}

func createContentWithData(
	hash hash.Hash,
	name string,
	data []byte,
) Content {
	return createContentWithDataAndAdditionAndRemoval(hash, name, data, nil, nil)
}

func createContentWithAddition(
	hash hash.Hash,
	name string,
	addition []hash.Hash,
) Content {
	return createContentWithDataAndAdditionAndRemoval(hash, name, nil, addition, nil)
}

func createContentWithRemoval(
	hash hash.Hash,
	name string,
	removal []hash.Hash,
) Content {
	return createContentWithDataAndAdditionAndRemoval(hash, name, nil, nil, removal)
}

func createContentWithDataAndAddition(
	hash hash.Hash,
	name string,
	data []byte,
	addition []hash.Hash,
) Content {
	return createContentWithDataAndAdditionAndRemoval(hash, name, data, addition, nil)
}

func createContentWithDataAndRemoval(
	hash hash.Hash,
	name string,
	data []byte,
	removal []hash.Hash,
) Content {
	return createContentWithDataAndAdditionAndRemoval(hash, name, data, nil, removal)
}

func createContentWithAdditionAndRemoval(
	hash hash.Hash,
	name string,
	addition []hash.Hash,
	removal []hash.Hash,
) Content {
	return createContentWithDataAndAdditionAndRemoval(hash, name, nil, addition, removal)
}

func createContentWithDataAndAdditionAndRemoval(
	hash hash.Hash,
	name string,
	data []byte,
	addition []hash.Hash,
	removal []hash.Hash,
) Content {
	return createContentWithDataAndAdditionAndRemoval(hash, name, data, addition, removal)
}

func createContentInternally(
	hash hash.Hash,
	name string,
	data []byte,
	addition []hash.Hash,
	removal []hash.Hash,
) Content {
	out := content{
		hash:     hash,
		name:     name,
		data:     data,
		addition: addition,
		removal:  removal,
	}

	return &out
}

// Hash returns the hash
func (obj *content) Hash() hash.Hash {
	return obj.hash
}

// Name returns the name
func (obj *content) Name() string {
	return obj.name
}

// HasData returns true if there is data, false otherwise
func (obj *content) HasData() bool {
	return obj.data != nil
}

// Data returns the data, if any
func (obj *content) Data() []byte {
	return obj.data
}

// HasWhitelistAddition returns true if there is a whitelist addition, false otherwise
func (obj *content) HasWhitelistAddition() bool {
	return obj.addition != nil
}

// WhitelistAddition returns the whitelist addition, if any
func (obj *content) WhitelistAddition() []hash.Hash {
	return obj.addition
}

// HasWhitelistRemoval returns true if there is a whitelist removal, false otherwise
func (obj *content) HasWhitelistRemoval() bool {
	return obj.removal != nil
}

// WhitelistRemoval returns the whitelist removal, if any
func (obj *content) WhitelistRemoval() []hash.Hash {
	return obj.removal
}
