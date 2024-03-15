package resources

import "github.com/steve-care-software/datastencil/domain/hash"

type native struct {
	hash    hash.Hash
	pSingle *uint8
	list    List
}

func createNativeWithSingle(
	hash hash.Hash,
	pSingle *uint8,
) Native {
	return createNativeInternally(hash, pSingle, nil)
}

func createNativeWithList(
	hash hash.Hash,
	list List,
) Native {
	return createNativeInternally(hash, nil, list)
}

func createNativeInternally(
	hash hash.Hash,
	pSingle *uint8,
	list List,
) Native {
	out := native{
		hash:    hash,
		pSingle: pSingle,
		list:    list,
	}

	return &out
}

// Hash returns the hash
func (obj *native) Hash() hash.Hash {
	return obj.hash
}

// IsSingle returns true if there is a single, flase otherwise
func (obj *native) IsSingle() bool {
	return obj.pSingle != nil
}

// Single returns the single, if any
func (obj *native) Single() *uint8 {
	return obj.pSingle
}

// IsList returns true if there is a list, flase otherwise
func (obj *native) IsList() bool {
	return obj.list != nil
}

// List returns the list, if any
func (obj *native) List() List {
	return obj.list
}
