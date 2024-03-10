package resources

type resource struct {
	name     string
	key      Field
	fields   Fields
	children Resources
}

func createResource(
	name string,
	key Field,
	fields Fields,
) Resource {
	return createResourceInternally(name, key, fields, nil)
}

func createResourceWithChildren(
	name string,
	key Field,
	fields Fields,
	children Resources,
) Resource {
	return createResourceInternally(name, key, fields, children)
}

func createResourceInternally(
	name string,
	key Field,
	fields Fields,
	children Resources,
) Resource {
	out := resource{
		name:     name,
		key:      key,
		fields:   fields,
		children: children,
	}

	return &out
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
