package resources

import "github.com/steve-care-software/datastencil/domain/hash"

type resource struct {
	hash     hash.Hash
	name     string
	key      Field
	fields   Fields
	children Resources
}

func createResource(
	hash hash.Hash,
	name string,
	key Field,
	fields Fields,
) Resource {
	return createResourceInternally(hash, name, key, fields, nil)
}

func createResourceWithChildren(
	hash hash.Hash,
	name string,
	key Field,
	fields Fields,
	children Resources,
) Resource {
	return createResourceInternally(hash, name, key, fields, children)
}

func createResourceInternally(
	hash hash.Hash,
	name string,
	key Field,
	fields Fields,
	children Resources,
) Resource {
	out := resource{
		hash:     hash,
		name:     name,
		key:      key,
		fields:   fields,
		children: children,
	}

	return &out
}

// Hash returns the hash
func (obj *resource) Hash() hash.Hash {
	return obj.hash
}

// Name returns the name
func (obj *resource) Name() string {
	return obj.name
}

// Key returns the key field
func (obj *resource) Key() Field {
	return obj.key
}

// Fields returns the fields
func (obj *resource) Fields() Fields {
	return obj.fields
}

// HasChildren returns true if there is children, false otherwise
func (obj *resource) HasChildren() bool {
	return obj.children != nil
}

// Children returns the children, if any
func (obj *resource) Children() Resources {
	return obj.children
}
