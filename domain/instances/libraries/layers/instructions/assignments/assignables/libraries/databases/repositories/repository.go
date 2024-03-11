package repositories

import "github.com/steve-care-software/datastencil/domain/hash"

type repository struct {
	hash       hash.Hash
	isSkeleton bool
	isHeight   bool
	list       string
	retrieve   string
}

func createRepositoryWithSkeleton(
	hash hash.Hash,
) Repository {
	return createRepositoryInternally(hash, true, false, "", "")
}

func createRepositoryWithHeight(
	hash hash.Hash,
) Repository {
	return createRepositoryInternally(hash, false, true, "", "")
}

func createRepositoryWithList(
	hash hash.Hash,
	list string,
) Repository {
	return createRepositoryInternally(hash, false, false, list, "")
}

func createRepositoryWithRetrieve(
	hash hash.Hash,
	retrieve string,
) Repository {
	return createRepositoryInternally(hash, false, false, "", retrieve)
}

func createRepositoryInternally(
	hash hash.Hash,
	isSkeleton bool,
	isHeight bool,
	list string,
	retrieve string,
) Repository {
	out := repository{
		hash:       hash,
		isSkeleton: isSkeleton,
		isHeight:   isHeight,
		list:       list,
		retrieve:   retrieve,
	}

	return &out
}

// Hash returns the hash, if any
func (obj *repository) Hash() hash.Hash {
	return obj.hash
}

// IsSkeleton returns true if there is a skeleton, false otherwise
func (obj *repository) IsSkeleton() bool {
	return obj.isSkeleton
}

// IsHeight returns true if there is a height, false otherwise
func (obj *repository) IsHeight() bool {
	return obj.isHeight
}

// IsList returns true if there is a list, false otherwise
func (obj *repository) IsList() bool {
	return obj.list != ""
}

// List returns the list, if any
func (obj *repository) List() string {
	return obj.list
}

// IsRetrieve returns true if there is a retrieve, false otherwise
func (obj *repository) IsRetrieve() bool {
	return obj.retrieve != ""
}

// Retrieve returns the retrieve, if any
func (obj *repository) Retrieve() string {
	return obj.retrieve
}
