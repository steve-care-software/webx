package resources

type resource struct {
	name         string
	relativePath string
}

func createResource(
	name string,
	relativePath string,
) Resource {
	out := resource{
		name:         name,
		relativePath: relativePath,
	}

	return &out
}

// Name returns the name
func (obj *resource) Name() string {
	return obj.name
}

// RelativePath returns the relative path
func (obj *resource) RelativePath() string {
	return obj.relativePath
}
